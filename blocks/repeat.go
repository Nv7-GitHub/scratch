package blocks

import "github.com/Nv7-Github/scratch/types"

type Repeat struct {
	*BasicBlock
	*Mouth

	Times types.Value
}

func (b *Blocks) NewRepeat(times types.Value) *Repeat {
	return &Repeat{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Mouth:      newMouth(),
		Times:      times,
	}
}

func (r *Repeat) Build() map[string]types.ScratchBlock {
	blk := r.BasicBlock.Build("control_repeat", map[string]types.ScratchInput{
		"TIMES":    r.Times.Build(),
		"SUBSTACK": types.NewScratchInputNoShadow(r.Blocks[0].ScratchID()),
	}, make(map[string]types.ScratchField))
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
	blk := r.BasicBlock.Build("control_repeat_until", map[string]types.ScratchInput{
		"CONDITION": r.Condition.Build(),
		"SUBSTACK":  types.NewScratchInputNoShadow(r.Blocks[0].ScratchID()),
	}, make(map[string]types.ScratchField))
	return r.Mouth.Build(blk, r.ScratchID())
}
