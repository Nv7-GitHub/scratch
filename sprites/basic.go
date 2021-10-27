package sprites

import (
	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

func newBasicSprite(name string) *BasicSprite {
	return &BasicSprite{
		Name:      name,
		Variables: make(map[string]*Variable),
		Lists:     make(map[string]*List),
		Costumes:  make([]*assets.Costume, 0),
		Sounds:    make([]*assets.Sound, 0),
	}
}

type BasicSprite struct {
	Name string

	Variables map[string]*Variable
	Lists     map[string]*List

	Costume  int
	Costumes []*assets.Costume

	Sounds []*assets.Sound
	Volume int

	comments map[string]string
}

func (b *BasicSprite) AddCostume(costume *assets.Costume) {
	b.Costumes = append(b.Costumes, costume)
}

func (b *BasicSprite) AddSound(sound *assets.Sound) {
	b.Sounds = append(b.Sounds, sound)
}

func (b *BasicSprite) Build() types.ScratchTargetBase {
	vars := make(map[string]types.ScratchVariableValue)
	for _, v := range b.Variables {
		vars[v.id] = types.ScratchVariableValue{v.Name, v.InitialValue}
	}

	lists := make(map[string]types.ScratchVariableValue)
	for _, v := range b.Lists {
		lists[v.id] = types.ScratchVariableValue{v.Name, v.InitialValues}
	}

	comments := make(map[string]types.ScratchComment)
	for id, c := range b.comments {
		comments[id] = types.ScratchComment{
			BlockID:   id,
			X:         0,
			Y:         0,
			Width:     0,
			Height:    0,
			Minimized: true,
			Text:      c,
		}
	}

	costumes := make([]types.ScratchCostume, len(b.Costumes))
	for i, costume := range b.Costumes {
		costumes[i] = costume.Build()
	}

	sounds := make([]types.ScratchSound, len(b.Sounds))
	for i, sound := range b.Sounds {
		sounds[i] = sound.Build()
	}

	return types.ScratchTargetBase{
		IsStage:    false,
		Name:       b.Name,
		Variables:  vars,
		Lists:      lists,
		Broadcasts: make(map[string]string), // Empty everywhere but stage
		// Blocks left for implementing sprite
		Comments:       comments,
		CurrentCostume: b.Costume,
		Costumes:       costumes,
		Sounds:         sounds,
		Volume:         b.Volume,
	}
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
