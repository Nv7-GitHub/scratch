package assets

import (
	"io"
)

type FS interface {
	Create(string) (io.Writer, error)
}

func Save(fs FS) error {
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
