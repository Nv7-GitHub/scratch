package blocks

import "github.com/Nv7-Github/scratch/types"

type Mouth struct {
	Blocks    []Block
	BlockVals []Block
}

func newMouth() *Mouth {
	return &Mouth{Blocks: make([]Block, 0)}
}

func (m *Mouth) Add(block Block) {
	_, ok := block.(BlockVal)
	if ok {
		m.BlockVals = append(m.BlockVals, block)
	} else {
		m.Blocks = append(m.Blocks, block)
	}
}

func (m *Mouth) Build(top types.ScratchBlock, topid string) map[string]types.ScratchBlock {
	out := make(map[string]types.ScratchBlock)
	for i, block := range m.Blocks {
		if i > 0 {
			block.SetPrevID(m.Blocks[i-1].ScratchID())
		} else {
			block.SetPrevID(topid)
		}
		if i < len(m.Blocks)-1 {
			block.SetNextID(m.Blocks[i+1].ScratchID())
		}

		simple, ok := block.(SimpleBlock)
		if ok {
			out[block.ScratchID()] = simple.Build()
		} else {
			// Mouth block
			blks := block.(MouthBlock).Build()
			for k, v := range blks {
				out[k] = v
			}
		}
	}
	out[topid] = top
	for _, val := range m.BlockVals {
		out[val.ScratchID()] = val.(SimpleBlock).Build()
	}
	return out
}
