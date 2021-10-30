package values

import "github.com/Nv7-Github/scratch/types"

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
