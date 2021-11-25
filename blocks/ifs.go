package blocks

import "github.com/Nv7-Github/scratch/types"

type IfBlock struct {
	*BasicBlock
	*Mouth

	Condition types.Value
}

func (i *IfBlock) Build() map[string]types.ScratchBlock {
	blk := i.BasicBlock.Build("control_if", map[string]types.ScratchInput{
		"CONDITION": i.Condition.Build(),
		"SUBSTACK":  types.NewScratchInputNoShadow(i.Blocks[0].ScratchID()),
	}, make(map[string]types.ScratchField))
	return i.Mouth.Build(blk, i.ScratchID())
}

func (b *Blocks) NewIf(condition types.Value) *IfBlock {
	return &IfBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Mouth:      newMouth(),
		Condition:  condition,
	}
}

type IfElseBlock struct {
	*BasicBlock
	*Mouth

	Condition types.Value
	Else      *Mouth
}

func (i *IfElseBlock) Build() map[string]types.ScratchBlock {
	blk := i.BasicBlock.Build("control_if_else", map[string]types.ScratchInput{
		"CONDITION": i.Condition.Build(),
		"SUBSTACK":  types.NewScratchInputNoShadow(i.Blocks[0].ScratchID()),
		"SUBSTACK2": types.NewScratchInputNoShadow(i.Else.Blocks[0].ScratchID()),
	}, make(map[string]types.ScratchField))
	stack := i.Mouth.Build(blk, i.ScratchID())

	// Combine stacks
	stack2 := i.Else.Build(blk, i.ScratchID())
	delete(stack2, i.ScratchID())
	for k, v := range stack2 {
		stack[k] = v
	}
	return stack
}

func (b *Blocks) NewIfElse(condition types.Value) *IfElseBlock {
	return &IfElseBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Mouth:      newMouth(),
		Condition:  condition,
		Else:       newMouth(),
	}
}
