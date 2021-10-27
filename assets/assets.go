package assets

import (
	"io"

	"github.com/Nv7-Github/scratch/types"
)

func Save(fs types.FS) error {
	for _, costume := range costumes {
		f, err := fs.Create(costume.filename())
		if err != nil {
			return err
		}

		_, err = io.Copy(f, costume.data)
		if err != nil {
			return err
		}
	}

	for _, sound := range sounds {
		f, err := fs.Create(sound.filename())
		if err != nil {
			return err
		}

		_, err = io.Copy(f, sound.data)
		if err != nil {
			return err
		}
	}

	return nil
}

func Clear() {
	costumes = make([]*Costume, 0)
	sounds = make([]*Sound, 0)
}
