package audio

import (
	"fmt"
	"log"
	"os"
	"server/utils/audio/wave"
)

type Mic struct {
}

type AudioFile struct {
	Path        string
	ContentType string
}

func NewMic() *Mic {
	return &Mic{}
}

type MicWave struct {
	params wave.WriterParam
}

func (ctx *Mic) Wave() *MicWave {
	return &MicWave{params: wave.WriterParam{
		WaveFormatType: 1,
		Channel:        1,
		SampleRate:     24000,
		BitsPerSample:  16,
	}}
}

// TODO: add setter e.g (PerSample, Rate)
func (ctx *MicWave) SetChannels(channels int) *MicWave {
	return ctx
}

func (ctx *MicWave) Record(path string, frames []byte) (*AudioFile, error) {
	path = fmt.Sprintf("%s.wav", path)

	w, err := os.Create(path)

	if err != nil {
		log.Fatalf("could not create file: %v", err)
	}

	defer w.Close()

	ctx.params.Out = w

	wavw, err := wave.NewWriter(ctx.params)

	if err != nil {
		return nil, err
	}

	defer wavw.Close()

	if _, err := wavw.Write(frames); err != nil {
		return nil, err
	}

	return &AudioFile{
		Path:        path,
		ContentType: "audio/wav",
	}, nil
}

// ----------- Python Package -----------
// channels=1, rate=24000, sample_width=2
// wf.setnchannels(channels)
// wf.setsampwidth(sample_width)
// wf.setframerate(rate)
// wf.writeframes(pcm)
// channels=1, rate=, sample_width=2

// ----------- GO PACKAGE -----------
// AudioFormat:   1,
// NumChans:      2,
// SampleRate:    44100,
// ByteRate:      176400,
// BlockAlign:    4,
// BitsPerSample: 16,
