package blocks

import "github.com/Nv7-Github/scratch/types"

type Block interface {
	ScratchID() string
	Build() types.ScratchBlock
}
