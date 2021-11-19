package blocks

import "github.com/Nv7-Github/scratch/types"

// SetVariable
type SetVariable struct {
	*BasicBlock

	variable *types.Variable
	val      types.Value
}

func (b *Blocks) NewSetVariable(variable *types.Variable, val types.Value) *SetVariable {
	return &SetVariable{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		variable:   variable,
		val:        val,
	}
}

func (s *SetVariable) Build() types.ScratchBlock {
	return s.BasicBlock.Build("data_setvariableto", map[string]types.ScratchInput{
		"VALUE": s.val.Build(),
	}, map[string]types.ScratchField{
		"VARIABLE": types.NewScratchValueFieldVariable(s.variable.Name, s.variable.ScratchID()),
	})
}

// ChangeVariable
type ChangeVariable struct {
	*BasicBlock

	variable *types.Variable
	val      types.Value
}

func (b *Blocks) NewChangeVariable(variable *types.Variable, val types.Value) *ChangeVariable {
	return &ChangeVariable{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		variable:   variable,
		val:        val,
	}
}

func (c *ChangeVariable) Build() types.ScratchBlock {
	return c.BasicBlock.Build("data_changevariableby", map[string]types.ScratchInput{
		"VALUE": c.val.Build(),
	}, map[string]types.ScratchField{
		"VARIABLE": types.NewScratchValueFieldVariable(c.variable.Name, c.variable.ScratchID()),
	})
}

// ChangeVariableVisibility
type ChangeVariableVisibility struct {
	*BasicBlock

	Show bool

	variable *types.Variable
}

func (b *Blocks) NewChangeVariableVisibility(variable *types.Variable, show bool) *ChangeVariableVisibility {
	return &ChangeVariableVisibility{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		variable:   variable,
		Show:       show,
	}
}

func (c *ChangeVariableVisibility) Build() types.ScratchBlock {
	blk := c.BasicBlock.Build("data_showvariable", make(map[string]types.ScratchInput), map[string]types.ScratchField{
		"VARIABLE": types.NewScratchValueFieldVariable(c.variable.Name, c.variable.ScratchID()),
	})
	if !c.Show {
		blk.Opcode = "data_hidevariable"
	}
	return blk
}
