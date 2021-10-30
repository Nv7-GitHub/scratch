package blocks

import "github.com/Nv7-Github/scratch/types"

type Repeat struct {
	*BasicBlock
	*Mouth

	Times int
}

func (b *Blocks) NewRepeat(times int) *Repeat {
	return &Repeat{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Mouth:      newMouth(),
		Times:      times,
	}
}

func (r *Repeat) Build() map[string]types.ScratchBlock {
	blk := types.ScratchBlock{
		Opcode: "control_repeat",
		Next:   r.next,
		Parent: r.prev,
		Inputs: map[string]types.ScratchInput{
			"TIMES":    types.NewScratchInputShadow(types.NewScratchInt(r.Times)),
			"SUBSTACK": types.NewScratchInputNoShadow(r.Blocks[0].ScratchID()),
		},
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: false,
	}
	return r.Mouth.Build(blk, r.ScratchID())
}

type RepeatUntil struct {
	*BasicBlock
	*Mouth

	Condition types.Value
}

func (b *Blocks) NewRepeatUntil(cond types.Value) *RepeatUntil {
	return &RepeatUntil{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Mouth:      newMouth(),
		Condition:  cond,
	}
}

func (r *RepeatUntil) Build() map[string]types.ScratchBlock {
	blk := types.ScratchBlock{
		Opcode: "control_repeat_until",
		Next:   r.next,
		Parent: r.prev,
		Inputs: map[string]types.ScratchInput{
			"CONDITION": r.Condition.Build(),
			"SUBSTACK":  types.NewScratchInputNoShadow(r.Blocks[0].ScratchID()),
		},
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: false,
	}
	return r.Mouth.Build(blk, r.ScratchID())
}
