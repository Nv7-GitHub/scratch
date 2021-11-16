package blocks

import "github.com/Nv7-Github/scratch/types"

type Blocks struct{}
type StageBlocks struct{ *Blocks }
type SpriteBlocks struct{ *Blocks }

func Clear() {

}

type Block interface {
	ScratchID() string
	SetNextID(string)
	SetPrevID(string)
}

type BlockVal interface {
	ScratchBlockVal()
}

type SimpleBlock interface {
	Block
	Build() types.ScratchBlock
}

type MouthBlock interface {
	Block
	Build() map[string]types.ScratchBlock
}

type BasicBlock struct {
	next *string
	prev *string
	id   string
}

func (b *BasicBlock) Build(opcode string, inputs map[string]types.ScratchInput, fields map[string]types.ScratchField) types.ScratchBlock {
	return types.ScratchBlock{
		Opcode:   opcode,
		Next:     b.next,
		Parent:   b.prev,
		Inputs:   inputs,
		Fields:   fields,
		Shadow:   false,
		TopLevel: false,
	}
}

func (b *BasicBlock) ScratchID() string {
	return b.id
}

func (b *BasicBlock) SetNextID(id string) {
	b.next = &id
}

func (b *BasicBlock) SetPrevID(id string) {
	b.prev = &id
}

func newBasicBlock(id string) *BasicBlock {
	return &BasicBlock{
		id: id,
	}
}
