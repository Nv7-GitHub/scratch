package sprites

type Variable struct {
	Name         string
	InitialValue interface{}

	id string
}

type List struct {
	Name          string
	InitialValues []interface{}

	id string
}
