package blocks

import "github.com/Nv7-Github/scratch/types"

type MathOperation int

const (
	MathOperationAdd MathOperation = iota
	MathOperationSubtract
	MathOperationMultiply
	MathOperationDivide
)

var mathOperationNames = map[MathOperation]string{
	MathOperationAdd:      "operator_add",
	MathOperationSubtract: "operator_subtract",
	MathOperationMultiply: "operator_multiply",
	MathOperationDivide:   "operator_divide",
}

type Math struct {
	*BasicBlock

	Operator MathOperation
	Num1     types.Value
	Num2     types.Value
}

func (m *Math) ScratchBlockVal() {}

func (m *Math) Build() types.ScratchBlock {
	return m.BasicBlock.Build(mathOperationNames[m.Operator], map[string]types.ScratchInput{
		"NUM1": m.Num1.Build(),
		"NUM2": m.Num2.Build(),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewMath(num1 types.Value, num2 types.Value, op MathOperation) *Math {
	return &Math{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Operator:   op,
		Num1:       num1,
		Num2:       num2,
	}
}
