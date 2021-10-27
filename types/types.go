package types

import "io"

type Value interface {
	Build() interface{}
}

type FS interface {
	Create(string) (io.Writer, error)
}
