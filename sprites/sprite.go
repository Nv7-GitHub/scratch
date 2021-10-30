package sprites

import (
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

type Sprite struct {
	*BasicSprite
	*blocks.SpriteStacks
	*blocks.SpriteBlocks

	Layer         int
	X, Y          int
	Size          int // in percent
	Direction     int // degrees
	Draggable     bool
	RotationStyle RotationStyle
	Visible       bool
}

func AddSprite(name string) *Sprite {
	s := &Sprite{
		BasicSprite:  newBasicSprite(name),
		SpriteStacks: &blocks.SpriteStacks{Stacks: blocks.NewStacks()},
		SpriteBlocks: &blocks.SpriteBlocks{},

		RotationStyle: RotationStyleAllAround,
		Layer:         len(Sprites) + 1,
		Direction:     90,
		Visible:       true,
		Size:          100,
	}
	Sprites = append(Sprites, s)
	return s
}

func (s *Sprite) Show() {
	s.Visible = true
}

func (s *Sprite) Hide() {
	s.Visible = false
}

func (s *Sprite) Build() types.ScratchTarget {
	basic := s.BasicSprite.Build()
	basic.Blocks = s.Stacks.Build()
	return &types.ScratchSprite{
		ScratchTargetBase: basic,

		LayerOrder:    s.Layer,
		X:             s.X,
		Y:             s.Y,
		Size:          s.Size,
		Direction:     s.Direction,
		Draggable:     s.Draggable,
		RotationStyle: rotationStyleNames[s.RotationStyle],
		Visible:       s.Visible,
	}
}

type RotationStyle int

const (
	RotationStyleAllAround RotationStyle = iota
	RotationStyleLeftRight
	RotationStyleNone
)

var rotationStyleNames = map[RotationStyle]string{
	RotationStyleAllAround: "all around",
	RotationStyleLeftRight: "left-right",
	RotationStyleNone:      "none",
}
