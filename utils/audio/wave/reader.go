package wave

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

type WaveReader interface {
	io.Reader
	io.Seeker
	io.ReaderAt
}

type Reader struct {
	input WaveReader

	size int64

	RiffChunk *RiffChunk
	FmtChunk  *FmtChunk
	DataChunk *DataReaderChunk

	originOfAudioData int64
	NumSamples        uint32
	ReadSampleNum     uint32
	SampleTime        int

	// LIST chunk
	extChunkSize int64
}

// Comment
func NewReader(fileName string) (*Reader, error) {
	fi, err := os.Stat(fileName)

	if err != nil {
		return &Reader{}, err
	}

	if fi.Size() > maxFileSize {
		return &Reader{}, fmt.Errorf("file is too large: %d bytes", fi.Size())
	}

	f, err := os.Open(fileName)

	if err != nil {
		return &Reader{}, err
	}

	defer f.Close()

	waveData, err := io.ReadAll(f)

	if err != nil {
		return &Reader{}, err
	}

	reader := new(Reader)
	reader.size = fi.Size()
	reader.input = bytes.NewReader(waveData)

	if err := reader.parseRiffChunk(); err != nil {
		panic(err)
	}

	if err := reader.parseFmtChunk(); err != nil {
		panic(err)
	}

	if err := reader.parseListChunk(); err != nil {
		panic(err)
	}

	if err := reader.parseDataChunk(); err != nil {
		panic(err)
	}

	reader.NumSamples = reader.DataChunk.Size / uint32(reader.FmtChunk.Data.BlockSize)
	reader.SampleTime = int(reader.NumSamples / reader.FmtChunk.Data.SamplesPerSec)

	return reader, nil
}

type csize struct {
	ChunkSize uint32
}

// Comment
func (rd *Reader) parseRiffChunk() error {
	chunkId := make([]byte, 4)

	if err := binary.Read(rd.input, binary.BigEndian, chunkId); err != nil {
		return err
	}

	if string(chunkId[:]) != riffChunkToken {
		return fmt.Errorf("file is not RIFF: %s", rd.RiffChunk.ID)
	}

	chunkSize := &csize{}

	if err := binary.Read(rd.input, binary.LittleEndian, chunkSize); err != nil {
		return err
	}

	if chunkSize.ChunkSize+8 != uint32(rd.size) {
		return fmt.Errorf("riff_chunk_size must be whole file size - 8bytes, expected(%d), actual(%d)", chunkSize.ChunkSize+8, rd.size)
	}

	format := make([]byte, 4)

	if err := binary.Read(rd.input, binary.BigEndian, format); err != nil {
		return err
	}

	if string(format[:]) != waveFormatType {
		return fmt.Errorf("file is not WAVE: %s", rd.RiffChunk.FormatType)
	}

	riffChunk := RiffChunk{
		ID:         chunkId,
		Size:       chunkSize.ChunkSize,
		FormatType: format,
	}

	rd.RiffChunk = &riffChunk

	return nil
}

// Comment
func (rd *Reader) parseFmtChunk() error {
	// TODO: remove os.SEEK_SET - (Deprecated)
	rd.input.Seek(riffChunkSize, os.SEEK_SET)

	chunkId := make([]byte, 4)
	err := binary.Read(rd.input, binary.BigEndian, chunkId)

	if err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	}
	if string(chunkId[:]) != fmtChunkToken {
		return fmt.Errorf("fmt chunk id must be \"%s\" but value is %s", fmtChunkToken, chunkId)
	}

	chunkSize := &csize{}
	err = binary.Read(rd.input, binary.LittleEndian, chunkSize)

	if err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	}

	if chunkSize.ChunkSize != fmtChunkSize {
		return fmt.Errorf("fmt chunk size must be %d but value is %d", fmtChunkSize, chunkSize.ChunkSize)
	}

	var fmtChunkData WavFmtChunkData

	if err = binary.Read(rd.input, binary.LittleEndian, &fmtChunkData); err != nil {
		return err
	}

	fmtChunk := FmtChunk{
		ID:   chunkId,
		Size: chunkSize.ChunkSize,
		Data: &fmtChunkData,
	}

	rd.FmtChunk = &fmtChunk

	return nil
}

// Comment
func (rd *Reader) parseListChunk() error {
	// TODO: remove os.SEEK_SET - (Deprecated)
	rd.input.Seek(listChunkOffset, os.SEEK_SET)

	chunkID := make([]byte, 4)

	if err := binary.Read(rd.input, binary.BigEndian, chunkID); err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	} else if string(chunkID[:]) != listChunkToken {
		return nil
	}

	chunkSize := make([]byte, 1)

	if err := binary.Read(rd.input, binary.LittleEndian, chunkSize); err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	}

	rd.extChunkSize = int64(chunkSize[0]) + 4 + 4

	return nil
}

// Comment
func (rd *Reader) getRiffChunkSizeOffset() int64 {
	return riffChunkSizeBaseOffset + rd.extChunkSize
}

// Comment
func (rd *Reader) parseDataChunk() error {
	// TODO: remove os.SEEK_SET - (Deprecated)
	originOfDataChunk, _ := rd.input.Seek(rd.getRiffChunkSizeOffset(), os.SEEK_SET)

	chunkId := make([]byte, 4)
	err := binary.Read(rd.input, binary.BigEndian, chunkId)

	if err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	}

	if string(chunkId[:]) != dataChunkToken {
		return fmt.Errorf("data chunk id must be \"%s\" but value is %s", dataChunkToken, chunkId)
	}

	chunkSize := &csize{}
	err = binary.Read(rd.input, binary.LittleEndian, chunkSize)

	if err == io.EOF {
		return fmt.Errorf("unexpected file end")
	} else if err != nil {
		return err
	}

	rd.originOfAudioData = originOfDataChunk + 8
	audioData := io.NewSectionReader(rd.input, rd.originOfAudioData, int64(chunkSize.ChunkSize))

	dataChunk := DataReaderChunk{
		ID:   chunkId,
		Size: chunkSize.ChunkSize,
		Data: audioData,
	}

	rd.DataChunk = &dataChunk

	return nil
}

// Comment
func (rd *Reader) Read(p []byte) (int, error) {
	return rd.DataChunk.Data.Read(p)
}

// Comment
func (rd *Reader) ReadRawSample() ([]byte, error) {
	size := rd.FmtChunk.Data.BlockSize
	sample := make([]byte, size)
	_, err := rd.Read(sample)

	// TODO: return error here
	if err == nil {
		rd.ReadSampleNum += 1
	}

	return sample, nil
}

// Comment
func (rd *Reader) ReadSample() ([]float64, error) {
	raw, err := rd.ReadRawSample()
	channel := int(rd.FmtChunk.Data.Channel)
	ret := make([]float64, channel)

	if err != nil {
		return ret, err
	}

	length := len(raw) / channel

	for i := 0; i < channel; i++ {
		tmp := bytesToInt(raw[length*i : length*(i+1)])

		switch rd.FmtChunk.Data.BitsPerSamples {
		case 8:
			ret[i] = float64(tmp-128) / 128.0
		case 16:
			ret[i] = float64(tmp) / 32768.0
		}

		if err != nil && err != io.EOF {
			return ret, err
		}
	}
	return ret, nil
}

// Comment
func (rd *Reader) ReadSampleInt() ([]int, error) {
	raw, err := rd.ReadRawSample()
	channels := int(rd.FmtChunk.Data.Channel)
	ret := make([]int, channels)
	length := len(raw) / channels

	if err != nil {
		return ret, err
	}

	for i := 0; i < channels; i++ {
		ret[i] = bytesToInt(raw[length*i : length*(i+1)])

		if err != nil && err != io.EOF {
			return ret, err
		}
	}

	return ret, nil
}

// Comment
func bytesToInt(b []byte) int {
	var ret int
	switch len(b) {
	case 1:
		// 0 ~ 128 ~ 255
		ret = int(b[0])
	case 2:
		// -32768 ~ 0 ~ 32767
		ret = int(b[0]) + int(b[1])<<8
	case 3:
		// HiReso / DVDAudio
		ret = int(b[0]) + int(b[1])<<8 + int(b[2])<<16
	default:
		ret = 0
	}
	return ret
}
