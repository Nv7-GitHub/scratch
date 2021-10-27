package blocks

import "github.com/Nv7-Github/scratch/types"

func Clear() {

}

type Block interface {
	ScratchID() string
	Build() types.ScratchBlock
}
