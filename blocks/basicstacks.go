package blocks

import "github.com/Nv7-Github/scratch/types"

type WhenFlagClicked struct {
	*BasicStack
}

func (w *WhenFlagClicked) Build(x, y int) map[string]types.ScratchBlock {
	id := types.GetRandomString()
	blk := types.ScratchBlock{
		Opcode:   "event_whenflagclicked",
		Inputs:   make(map[string]types.ScratchInput),
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: true,
		X:        &x,
		Y:        &y,
	}
	b := w.BasicStack.Build(blk, id)
	return b
}

func (s *Stacks) NewWhenFlagClicked() *WhenFlagClicked {
	return &WhenFlagClicked{newBasicStack()}
}
