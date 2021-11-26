package blocks

import "github.com/Nv7-Github/scratch/types"

type Concat struct {
	*BasicBlock

	Val1 types.Value
	Val2 types.Value
}

func (c *Concat) ScratchBlockVal() {}

func (c *Concat) Build() types.ScratchBlock {
	return c.BasicBlock.Build("operator_join", map[string]types.ScratchInput{
		"STRING1": c.Val1.Build(),
		"STRING2": c.Val2.Build(),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewConcat(val1, val2 types.Value) *Concat {
	return &Concat{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Val1:       val1,
		Val2:       val2,
	}
}

type StringIndex struct {
	*BasicBlock

	Val   types.Value
	Index types.Value
}

func (s *StringIndex) ScratchBlockVal() {}

func (s *StringIndex) Build() types.ScratchBlock {
	return s.BasicBlock.Build("operator_letter_of", map[string]types.ScratchInput{
		"STRING": s.Val.Build(),
		"LETTER": s.Index.Build(),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewStringIndex(val, index types.Value) *StringIndex {
	return &StringIndex{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		Val:   val,
		Index: index,
	}
}

type StringLength struct {
	*BasicBlock

	Val types.Value
}

func (s *StringLength) ScratchBlockVal() {}

func (s *StringLength) Build() types.ScratchBlock {
	return s.BasicBlock.Build("operator_length", map[string]types.ScratchInput{
		"STRING": s.Val.Build(),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewStringLength(val types.Value) *StringLength {
	return &StringLength{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		Val: val,
	}
}

type StringContains struct {
	*BasicBlock

	Val    types.Value
	Substr types.Value
}

func (s *StringContains) ScratchBlockVal() {}

func (s *StringContains) Build() types.ScratchBlock {
	return s.BasicBlock.Build("operator_contains", map[string]types.ScratchInput{
		"STRING1": s.Val.Build(),
		"STRING2": s.Substr.Build(),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewStringContains(val, substr types.Value) *StringContains {
	return &StringContains{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		Val:    val,
		Substr: substr,
	}
}
