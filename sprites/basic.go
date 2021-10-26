package sprites

type BasicSprite struct {
	Name         string
	Variables    map[string]*Variable
	Broadcasts   []string
	broadcastIDs []string

	Costume  int
	Costumes []*Costume
}

type Variable struct {
	Name         string
	InitialValue interface{}

	id string
}

type Costume struct {
	id string

	RotationCenterX int
	RotationCenterY int
}

type Sound struct {
	id          string
	rate        int
	samplecount int
}
