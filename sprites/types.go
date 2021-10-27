package sprites

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
