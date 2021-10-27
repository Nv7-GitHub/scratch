package assets

import (
	"io"

	"github.com/Nv7-Github/scratch/types"
)

var sounds = make([]*Sound, 0)

type Sound struct {
	Name string

	id          string
	rate        int
	sampleCount int
	data        io.Reader
}

func (s *Sound) Filename() string {
	return s.id + "." + s.Format()
}

func (s *Sound) Format() string {
	return "wav"
}

func (s *Sound) Rate() int {
	return s.rate
}

func (s *Sound) SampleCount() int {
	return s.sampleCount
}

func (s *Sound) ScratchID() string {
	return s.id
}

func GetSound(name string, data io.Reader, rate int, sampleCount int) *Sound {
	sound := &Sound{
		Name:        name,
		data:        data,
		rate:        rate,
		sampleCount: sampleCount,

		id: types.GetRandomString(),
	}

	sounds = append(sounds, sound)

	return sound
}
