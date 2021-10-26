package sprites

import "github.com/Nv7-Github/scratch/blocks"

type BasicSprite struct {
	Name string

	Variables map[string]*Variable
	Lists     map[string]*List

	Broadcasts   []string
	broadcastIDs []string

	Costume  int
	Costumes []*Costume

	Sounds []*Sound
	Volume int

	Blocks []blocks.Block

	comments map[string]string
}

func (b *BasicSprite) GetComment(block blocks.Block) string {
	return b.comments[block.ScratchID()]
}

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
