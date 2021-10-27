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

func (s *Sound) filename() string {
	return s.id + ".wav"
}

func (s *Sound) Build() types.ScratchSound {
	return types.ScratchSound{
		ScratchAsset: types.ScratchAsset{
			AssetID:    s.id,
			Name:       s.Name,
			Md5Ext:     s.filename(),
			DataFormat: "wav",
		},
		Format:      "",
		Rate:        s.rate,
		SampleCount: s.sampleCount,
	}
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
