package sprites

import "github.com/Nv7-Github/scratch/types"

var Sprites []*Sprite

func Clear() {
	Sprites = make([]*Sprite, 0)
}

func Build() []types.ScratchTarget {
	out := make([]types.ScratchTarget, len(Sprites))
	for i, sprite := range Sprites {
		out[i] = sprite.Build()
	}
	return out
}
