package sprites

import "github.com/Nv7-Github/scratch/types"

type Sprite interface {
	Build() types.ScratchTarget
}
