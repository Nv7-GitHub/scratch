package blocks

import "github.com/Nv7-Github/scratch/types"

type WhenBroadcastReceived struct {
	*BasicStack

	message *types.BroadcastMessage
}

func (w *WhenBroadcastReceived) Build() map[string]types.ScratchBlock {
	id := types.GetRandomString()
	blk := types.ScratchBlock{
		Opcode: "event_whenbroadcastreceived",
		Inputs: make(map[string]types.ScratchInput),
		Fields: map[string]types.ScratchField{
			"BROADCAST_OPTION": types.NewScratchFieldBroadcast(w.message.Name, w.message.ScratchID()),
		},
		Shadow:   false,
		TopLevel: true,
		X:        &w.X,
		Y:        &w.Y,
	}
	b := w.BasicStack.Build(blk, id)
	return b
}

func (s *Stacks) NewBroadcastReceive(msg *types.BroadcastMessage) Stack {
	return s.addStack(&WhenBroadcastReceived{newBasicStack(), msg})
}

type BroadcastBlock struct {
	*BasicBlock

	message *types.BroadcastMessage
}

func (b *BroadcastBlock) Build() types.ScratchBlock {
	return b.BasicBlock.Build("event_broadcast", map[string]types.ScratchInput{
		"BROADCAST_INPUT": types.NewScratchInputShadow(types.NewScratchBroadcast(b.message.Name, b.message.ScratchID())),
	}, make(map[string]types.ScratchField))
}

func (b *Blocks) NewBroadcast(msg *types.BroadcastMessage) Block {
	return &BroadcastBlock{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		message: msg,
	}
}
