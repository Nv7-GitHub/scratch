package blocks

import "github.com/Nv7-Github/scratch/types"

type CompareOperand int

const (
	CompareEqual CompareOperand = iota
	CompareLessThan
	CompareGreaterThan
)

var compareOperandNames = map[CompareOperand]string{
	CompareEqual:       "operator_equals",
	CompareLessThan:    "operator_lt",
	CompareGreaterThan: "operator_gt",
}

type Compare struct {
	*BasicBlock

	Op   CompareOperand
	Val1 types.Value
	Val2 types.Value
}

func (c *Compare) Build() types.ScratchBlock {
	return types.ScratchBlock{
		Opcode: compareOperandNames[c.Op],
		Next:   c.next,
		Parent: c.prev,
		Inputs: map[string]types.ScratchInput{
			"OPERAND1": types.NewScratchInputShadow(c.Val1.Build()),
			"OPERAND2": types.NewScratchInputShadow(c.Val2.Build()),
		},
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: false,
	}
}

func (b *Blocks) NewCompare(val1 types.Value, val2 types.Value, op CompareOperand) *Compare {
	return &Compare{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		Op:         op,
		Val1:       val1,
		Val2:       val2,
	}
}
