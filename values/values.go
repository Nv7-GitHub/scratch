package values

import (
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

type IntValue struct {
	Value int
}

func NewIntValue(val int) types.Value {
	return &IntValue{Value: val}
}

func (i *IntValue) Build() types.ScratchInput {
	return types.NewScratchInputShadow(types.NewScratchInt(i.Value))
}

type FloatValue struct {
	Value float64
}

func (f *FloatValue) Build() types.ScratchInput {
	return types.NewScratchInputShadow(types.NewScratchFloat(f.Value))
}

func NewFloatValue(val float64) types.Value {
	return &FloatValue{Value: val}
}

type StringValue struct {
	Value string
}

func NewStringValue(value string) types.Value {
	return &StringValue{Value: value}
}

func (s *StringValue) Build() types.ScratchInput {
	return types.NewScratchInputShadow(types.NewScratchString(s.Value))
}

type VariableValue struct {
	Variable *types.Variable
}

func NewVariableValue(variable *types.Variable) types.Value {
	return &VariableValue{Variable: variable}
}

func (v *VariableValue) Build() types.ScratchInput {
	return types.NewScratchInputObscured(types.NewScratchVariable(v.Variable.Name, v.Variable.ScratchID()), types.NewScratchString(""))
}

type ListValue struct {
	List *types.List
}

func NewListValue(list *types.List) types.Value {
	return &ListValue{List: list}
}

func (l *ListValue) Build() types.ScratchInput {
	return types.NewScratchInputObscured(types.NewScratchList(l.List.Name, l.List.ScratchID()), types.NewScratchString(""))
}

type BlockValue struct {
	Block blocks.Block
}

func NewBlockValue(block blocks.BlockVal) types.Value {
	return &BlockValue{Block: block}
}

func (b *BlockValue) Build() types.ScratchInput {
	return types.NewScratchInputShadow(types.NewScratchBlockInput(b.Block.ScratchID()))
}
