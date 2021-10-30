package blocks

import "github.com/Nv7-Github/scratch/types"

type WhenFlagClicked struct {
	*BasicStack
}

func (w *WhenFlagClicked) Build() map[string]types.ScratchBlock {
	id := types.GetRandomString()
	blk := types.ScratchBlock{
		Opcode:   "event_whenflagclicked",
		Inputs:   make(map[string]types.ScratchInput),
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: true,
		X:        &w.X,
		Y:        &w.Y,
	}
	b := w.BasicStack.Build(blk, id)
	return b
}

func (s *Stacks) NewWhenFlagClicked() Stack {
	return s.addStack(&WhenFlagClicked{newBasicStack()})
}
