package sprites

import (
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

func newBasicSprite(name string) *BasicSprite {
	return &BasicSprite{
		Name:      name,
		Variables: make(map[string]*Variable),
		Lists:     make(map[string]*List),
		Costumes:  make([]*Costume, 0),
		Sounds:    make([]*Sound, 0),
	}
}

type BasicSprite struct {
	Name string

	Variables map[string]*Variable
	Lists     map[string]*List

	Costume  int
	Costumes []*Costume

	Sounds []*Sound
	Volume int

	comments map[string]string
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
		costumes[i] = types.ScratchCostume{
			ScratchAsset: types.ScratchAsset{
				AssetID:    costume.id,
				Name:       costume.Name,
				Md5Ext:     costume.id + costume.extension,
				DataFormat: costume.extension,
			},
			BitmapResolution: 1,
			RotationCenterX:  costume.RotationCenterX,
			RotationCenterY:  costume.RotationCenterY,
		}
	}

	sounds := make([]types.ScratchSound, len(b.Sounds))
	for i, sound := range b.Sounds {
		sounds[i] = types.ScratchSound{
			ScratchAsset: types.ScratchAsset{
				AssetID:    sound.id,
				Name:       sound.Name,
				Md5Ext:     sound.id + "wav",
				DataFormat: "wav",
			},
			Format:      "",
			Rate:        sound.rate,
			SampleCount: sound.sampleCount,
		}
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

type Costume struct {
	id        string
	extension string

	Name            string
	RotationCenterX int
	RotationCenterY int
}

type Sound struct {
	Name string

	id          string
	rate        int
	sampleCount int
}
