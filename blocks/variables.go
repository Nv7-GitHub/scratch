package blocks

import "github.com/Nv7-Github/scratch/types"

type SetVariable struct {
	*BasicBlock

	variable *types.Variable
	val      types.Value
	Change   bool
}

func (b *Blocks) NewSetVariable(variable *types.Variable, val types.Value) *SetVariable {
	return &SetVariable{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		variable:   variable,
		val:        val,
	}
}

func (s *SetVariable) Build() types.ScratchBlock {
	blk := types.ScratchBlock{
		Opcode:   "data_setvariableto",
		Next:     s.next,
		Parent:   s.prev,
		Inputs:   map[string]types.ScratchInput{"VALUE": s.val.Build()},
		Fields:   map[string]types.ScratchField{"VARIABLE": types.NewScratchValueFieldVariable(s.variable.Name, s.variable.ScratchID())},
		Shadow:   false,
		TopLevel: false,
	}
	if s.Change {
		blk.Opcode = "data_changevariableby"
	}
	return blk
}
