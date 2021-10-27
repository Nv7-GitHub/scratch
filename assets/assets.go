package assets

import (
	"io"

	"github.com/Nv7-Github/scratch/types"
)

func Save(fs types.FS) error {
	for _, costume := range costumes {
		f, err := fs.Create(costume.Filename())
		if err != nil {
			return err
		}

		_, err = io.Copy(f, costume.data)
		if err != nil {
			return err
		}
	}

	for _, sound := range sounds {
		f, err := fs.Create(sound.Filename())
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
