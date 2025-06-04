package wave

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
)

const (
	maxFileSize             = 2 << 31
	riffChunkSize           = 12
	listChunkOffset         = 36
	riffChunkSizeBaseOffset = 36 // RIFFChunk(12byte) + fmtChunk(24byte) = 36byte
	fmtChunkSize            = 16
)

var (
	riffChunkToken = "RIFF"
	waveFormatType = "WAVE"
	fmtChunkToken  = "fmt "
	listChunkToken = "LIST"
	dataChunkToken = "data"
)

// 12byte
type RiffChunk struct {
	ID         []byte // 'RIFF'
	Size       uint32 // 36bytes + data_chunk_size or whole_file_size - 'RIFF'+ChunkSize (8byte)
	FormatType []byte // 'WAVE'
}

// 8 + 16 = 24byte
type FmtChunk struct {
	ID   []byte // 'fmt '
	Size uint32 // 16
	Data *WavFmtChunkData
}

// 16byte
type WavFmtChunkData struct {
	WaveFormatType uint16
	Channel        uint16
	SamplesPerSec  uint32
	BytesPerSec    uint32
	BlockSize      uint16
	BitsPerSamples uint16
}

// data 読み込み
type DataReader interface {
	io.Reader
	io.ReaderAt
}

type DataReaderChunk struct {
	ID   []byte
	Size uint32
	Data DataReader
}

type DataWriterChunk struct {
	ID   []byte
	Size uint32
	Data *bytes.Buffer
}

type WriterParam struct {
	Out            io.WriteCloser
	WaveFormatType int
	Channel        int
	SampleRate     int
	BitsPerSample  int
}

type Writer struct {
	out            io.WriteCloser
	writtenSamples int

	riffChunk *RiffChunk
	fmtChunk  *FmtChunk
	dataChunk *DataWriterChunk
}

// Comment
func NewWriter(param WriterParam) (*Writer, error) {
	w := &Writer{}
	w.out = param.Out

	blockSize := uint16(param.BitsPerSample*param.Channel) / 8
	samplesPerSec := uint32(int(blockSize) * param.SampleRate)

	// riff chunk
	w.riffChunk = &RiffChunk{
		ID:         []byte(riffChunkToken),
		FormatType: []byte(waveFormatType),
	}

	// fmt chunk
	w.fmtChunk = &FmtChunk{
		ID:   []byte(fmtChunkToken),
		Size: uint32(fmtChunkSize),
	}

	w.fmtChunk.Data = &WavFmtChunkData{
		WaveFormatType: uint16(param.WaveFormatType),
		Channel:        uint16(param.Channel),
		SamplesPerSec:  uint32(param.SampleRate),
		BytesPerSec:    samplesPerSec,
		BlockSize:      uint16(blockSize),
		BitsPerSamples: uint16(param.BitsPerSample),
	}

	// data chunk
	w.dataChunk = &DataWriterChunk{
		ID:   []byte(dataChunkToken),
		Data: bytes.NewBuffer([]byte{}),
	}

	return w, nil
}

// Comment
func (w *Writer) WriteInt8(samples []uint8) (int, error) {
	buf := new(bytes.Buffer)

	for i := 0; i < len(samples); i++ {
		err := binary.Write(buf, binary.LittleEndian, samples[i])

		if err != nil {
			return 0, err
		}
	}

	n, err := w.Write(buf.Bytes())

	return n, err
}

// Comment
func (w *Writer) WriteInt16(samples []int16) (int, error) {
	buf := new(bytes.Buffer)

	for i := 0; i < len(samples); i++ {
		err := binary.Write(buf, binary.LittleEndian, samples[i])

		if err != nil {
			return 0, err
		}
	}

	n, err := w.Write(buf.Bytes())

	return n, err
}

// Comment
func (w *Writer) Write(p []byte) (int, error) {
	blockSize := int(w.fmtChunk.Data.BlockSize)

	if len(p) < blockSize {
		return 0, fmt.Errorf("writing data need at least %d bytes", blockSize)
	}

	if len(p)%blockSize != 0 {
		return 0, fmt.Errorf("writing data must be a multiple of %d bytes", blockSize)
	}

	num := len(p) / blockSize

	n, err := w.dataChunk.Data.Write(p)

	if err == nil {
		w.writtenSamples += num
	}

	return n, err
}

type errWriter struct {
	w   io.Writer
	err error
}

// Comment
func (ew *errWriter) Write(order binary.ByteOrder, data interface{}) {
	if ew.err != nil {
		return
	}
	ew.err = binary.Write(ew.w, order, data)
}

// Comment
func (w *Writer) Close() error {
	data := w.dataChunk.Data.Bytes()
	dataSize := uint32(len(data))
	w.riffChunk.Size = uint32(len(w.riffChunk.ID)) + (8 + w.fmtChunk.Size) + (8 + dataSize)
	w.dataChunk.Size = dataSize

	ew := &errWriter{w: w.out}

	// riff chunk
	ew.Write(binary.BigEndian, w.riffChunk.ID)
	ew.Write(binary.LittleEndian, w.riffChunk.Size)
	ew.Write(binary.BigEndian, w.riffChunk.FormatType)

	// fmt chunk
	ew.Write(binary.BigEndian, w.fmtChunk.ID)
	ew.Write(binary.LittleEndian, w.fmtChunk.Size)
	ew.Write(binary.LittleEndian, w.fmtChunk.Data)

	//data chunk
	ew.Write(binary.BigEndian, w.dataChunk.ID)
	ew.Write(binary.LittleEndian, w.dataChunk.Size)

	if ew.err != nil {
		return ew.err
	}

	if _, err := w.out.Write(data); err != nil {
		return err
	}

	if err := w.out.Close(); err != nil {
		return err
	}

	return nil
}
