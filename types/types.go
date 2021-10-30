package types

import "io"

type Value interface {
	Build() ScratchInput
}

type FS interface {
	Create(string) (io.Writer, error)
}

type Variable struct {
	Name         string
	InitialValue interface{}

	Local bool

	id         string
	spriteName string
}

func (v *Variable) ScratchID() string {
	return v.id
}

func (v *Variable) SpriteName() string {
	return v.spriteName
}

func (v *Variable) SetScratchID(id string) {
	v.id = id
}

func (v *Variable) SetScratchSpriteName(name string) {
	v.spriteName = name
}

type List struct {
	Name          string
	InitialValues []interface{}

	Local bool

	id         string
	spriteName string
}

func (l *List) ScratchID() string {
	return l.id
}

func (l *List) SpriteName() string {
	return l.spriteName
}

func (l *List) SetScratchListID(id string) {
	l.id = id
}

func (l *List) SetScratchSpriteName(name string) {
	l.spriteName = name
}
