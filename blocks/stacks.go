package blocks

import "github.com/Nv7-Github/scratch/types"

type Stack interface {
	Build(x, y int) map[string]types.ScratchBlock
	Add(Block)
}

type Stacks struct{}
type StageStacks struct{ *Stacks }
type SpriteStacks struct{ *Stacks }

type BasicStack struct {
	blocks []Block
}

func (b *BasicStack) Add(block Block) {
	b.blocks = append(b.blocks, block)
}

func (b *BasicStack) Build(top types.ScratchBlock, topid string) map[string]types.ScratchBlock {
	blocks := make(map[string]types.ScratchBlock)
	for i, block := range b.blocks {
		if i > 0 {
			block.SetPrevID(b.blocks[i-1].ScratchID())
		} else {
			id := block.ScratchID()
			top.Next = &id
			block.SetPrevID(topid)
		}
		if i < len(b.blocks)-1 {
			block.SetNextID(b.blocks[i+1].ScratchID())
		}
		blocks[block.ScratchID()] = block.Build()
	}
	blocks[topid] = top
	return blocks
}

func newBasicStack() *BasicStack {
	return &BasicStack{
		blocks: make([]Block, 0),
	}
}
