package blocks

import "github.com/Nv7-Github/scratch/types"

// AddToList
type AddToList struct {
	*BasicBlock

	List  *types.List
	Value types.Value
}

func (a *AddToList) Build() types.ScratchBlock {
	return a.BasicBlock.Build("data_addtolist", map[string]types.ScratchInput{
		"ITEM": a.Value.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(a.List.Name, a.List.ScratchID()),
	})
}

func (b Blocks) NewAddToList(list *types.List, value types.Value) *AddToList {
	return &AddToList{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		List:       list,
		Value:      value,
	}
}

// DeleteFromList
type DeleteFromList struct {
	*BasicBlock

	List  *types.List
	Index types.Value
}

func (b *DeleteFromList) Build() types.ScratchBlock {
	return b.BasicBlock.Build("data_deleteoflist", map[string]types.ScratchInput{
		"INDEX": b.Index.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(b.List.Name, b.List.ScratchID()),
	})
}

func (b *Blocks) NewDeleteFromList(list *types.List, index types.Value) *DeleteFromList {
	return &DeleteFromList{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		List:       list,
		Index:      index,
	}
}

// DeleteAllFromList
type DeleteAllFromList struct {
	*BasicBlock

	List *types.List
}

func (d *DeleteAllFromList) Build() types.ScratchBlock {
	return d.BasicBlock.Build("data_deletealloflist", make(map[string]types.ScratchInput), map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(d.List.Name, d.List.ScratchID()),
	})
}

func (b *Blocks) NewDeleteAllFromList(list *types.List) *DeleteAllFromList {
	return &DeleteAllFromList{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List: list,
	}
}
