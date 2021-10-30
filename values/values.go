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

func (i *IntValue) Build() types.ScratchValue {
	return types.NewScratchInt(i.Value)
}

type FloatValue struct {
	Value float64
}

func (f *FloatValue) Build() types.ScratchValue {
	return types.NewScratchFloat(f.Value)
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

func (s *StringValue) Build() types.ScratchValue {
	return types.NewScratchString(s.Value)
}

type VariableValue struct {
	Variable *types.Variable
}

func NewVariableValue(variable *types.Variable) types.Value {
	return &VariableValue{Variable: variable}
}

func (v *VariableValue) Build() types.ScratchValue {
	return types.NewScratchVariable(v.Variable.Name, v.Variable.ScratchID())
}

type BlockValue struct {
	Block blocks.Block
}

func NewBlockValue(block blocks.Block) types.Value {
	return &BlockValue{Block: block}
}

func (b *BlockValue) Build() types.ScratchValue {
	return types.NewScratchBlockInput(b.Block.ScratchID())
}
