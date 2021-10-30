package blocks

import "github.com/Nv7-Github/scratch/types"

type Stack interface {
	Build() map[string]types.ScratchBlock
	Add(Block)
}

type Stacks struct {
	stacks []Stack
}

func NewStacks() *Stacks {
	return &Stacks{
		stacks: make([]Stack, 0),
	}
}

func (s *Stacks) Build() map[string]types.ScratchBlock {
	out := make(map[string]types.ScratchBlock)
	for _, stack := range s.stacks {
		built := stack.Build()
		for id, block := range built {
			out[id] = block
		}
	}
	return out
}

func (s *Stacks) addStack(stack Stack) Stack {
	s.stacks = append(s.stacks, stack)
	return stack
}

type StageStacks struct{ *Stacks }
type SpriteStacks struct{ *Stacks }

type BasicStack struct {
	blocks []Block
	X      int
	Y      int
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

		simple, ok := block.(SimpleBlock)
		if ok {
			blocks[block.ScratchID()] = simple.Build()
		} else {
			// Mouth block
			blks := block.(MouthBlock).Build()
			for k, v := range blks {
				blocks[k] = v
			}
		}
	}
	blocks[topid] = top
	return blocks
}

func newBasicStack() *BasicStack {
	return &BasicStack{
		blocks: make([]Block, 0),
	}
}
