package blocks

import "github.com/Nv7-Github/scratch/types"

type SayBlock struct {
	*BasicBlock

	time   bool
	length float64
	val    types.Value
}

func (s *SayBlock) Build() types.ScratchBlock {
	blk := s.BasicBlock.Build("looks_say", map[string]types.ScratchInput{
		"MESSAGE": s.val.Build(),
	}, make(map[string]types.ScratchField))
	if s.time {
		blk.Opcode = "looks_sayforsecs"
		blk.Inputs["SECS"] = types.NewScratchInputShadow(types.NewScratchFloat(s.length))
	}
	return blk
}

func (b *Blocks) NewSayBlock(text types.Value) *SayBlock {
	return &SayBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		time:       false,
		length:     0,
		val:        text,
	}
}

func (b *Blocks) NewSayForTimeBlock(text types.Value, time float64) *SayBlock {
	return &SayBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		time:       true,
		length:     time,
		val:        text,
	}
}
