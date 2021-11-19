package blocks

import (
	"github.com/Nv7-Github/scratch/types"
)

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

func (b *Blocks) NewAddToList(list *types.List, value types.Value) *AddToList {
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

// InsertInList
type InsertInList struct {
	*BasicBlock

	List  *types.List
	Index types.Value
	Value types.Value
}

func (i *InsertInList) Build() types.ScratchBlock {
	return i.BasicBlock.Build("data_insertatlist", map[string]types.ScratchInput{
		"ITEM":  i.Value.Build(),
		"INDEX": i.Index.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(i.List.Name, i.List.ScratchID()),
	})
}

func (b *Blocks) NewInsertInList(list *types.List, index types.Value, value types.Value) *InsertInList {
	return &InsertInList{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List:  list,
		Index: index,
		Value: value,
	}
}

// ReplaceInList
type ReplaceInList struct {
	*BasicBlock

	List  *types.List
	Index types.Value
	Value types.Value
}

func (r *ReplaceInList) Build() types.ScratchBlock {
	return r.BasicBlock.Build("data_replaceitemoflist", map[string]types.ScratchInput{
		"INDEX": r.Index.Build(),
		"ITEM":  r.Value.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(r.List.Name, r.List.ScratchID()),
	})
}

func (b *Blocks) NewReplaceInList(list *types.List, index types.Value, value types.Value) *ReplaceInList {
	return &ReplaceInList{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List:  list,
		Index: index,
		Value: value,
	}
}

// ItemOfList
type ItemOfList struct {
	*BasicBlock

	List  *types.List
	Index types.Value
}

func (i *ItemOfList) Build() types.ScratchBlock {
	return i.BasicBlock.Build("data_itemoflist", map[string]types.ScratchInput{
		"INDEX": i.Index.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(i.List.Name, i.List.ScratchID()),
	})
}

func (i *ItemOfList) ScratchBlockVal() {}

func (b *Blocks) NewItemOfList(list *types.List, index types.Value) *ItemOfList {
	return &ItemOfList{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List:  list,
		Index: index,
	}
}

// FindInList
type FindInList struct {
	*BasicBlock

	List  *types.List
	Value types.Value
}

func (f *FindInList) ScratchBlockVal() {}

func (f *FindInList) Build() types.ScratchBlock {
	return f.BasicBlock.Build("data_itemnumoflist", map[string]types.ScratchInput{
		"ITEM": f.Value.Build(),
	}, map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(f.List.Name, f.List.ScratchID()),
	})
}

// LengthOfList
type LengthOfList struct {
	*BasicBlock

	List *types.List
}

func (l *LengthOfList) ScratchBlockVal() {}

func (l *LengthOfList) Build() types.ScratchBlock {
	return l.BasicBlock.Build("data_lengthoflist", make(map[string]types.ScratchInput), map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(l.List.Name, l.List.ScratchID()),
	})
}

func (b *Blocks) NewLengthOfList(list *types.List) *LengthOfList {
	return &LengthOfList{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List: list,
	}
}

// ChangeListVisibility
type ChangeListVisibility struct {
	*BasicBlock

	List *types.List
	Show bool
}

func (c *ChangeListVisibility) Build() types.ScratchBlock {
	blk := c.BasicBlock.Build("data_showlist", make(map[string]types.ScratchInput), map[string]types.ScratchField{
		"LIST": types.NewScratchValueFieldVariable(c.List.Name, c.List.ScratchID()),
	})
	if !c.Show {
		blk.Opcode = "data_hidelist"
	}
	return blk
}

func (b *Blocks) NewChangeListVisibility(list *types.List, show bool) *ChangeListVisibility {
	return &ChangeListVisibility{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		List: list,
		Show: show,
	}
}
