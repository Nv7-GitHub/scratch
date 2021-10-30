package blocks

import "github.com/Nv7-Github/scratch/types"

type SayBlock struct {
	*BasicBlock

	time   bool
	length float64
	text   string
}

func (s *SayBlock) Build() types.ScratchBlock {
	blk := types.ScratchBlock{
		Opcode: "looks_say",
		Next:   s.next,
		Parent: s.prev,
		Inputs: map[string]types.ScratchInput{
			"MESSAGE": types.NewScratchInputShadow(types.NewScratchString(s.text)),
		},
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: false,
	}
	if s.time {
		blk.Opcode = "looks_sayforsecs"
		blk.Inputs["SECS"] = types.NewScratchInputShadow(types.NewScratchFloat(s.length))
	}
	return blk
}

func (b *Blocks) NewSayBlock(text string) *SayBlock {
	return &SayBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		time:       false,
		length:     0,
		text:       text,
	}
}

func (b *Blocks) NewSayForTimeBlock(text string, time float64) *SayBlock {
	return &SayBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		time:       true,
		length:     time,
		text:       text,
	}
}
