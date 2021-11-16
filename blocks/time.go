package blocks

import "github.com/Nv7-Github/scratch/types"

type Wait struct {
	*BasicBlock

	time types.Value
}

func (w *Wait) Build() types.ScratchBlock {
	return w.BasicBlock.Build("control_wait", map[string]types.ScratchInput{
		"DURATION": w.time.Build(),
	}, make(map[string]types.ScratchField))
}

type WaitUntil struct {
	*BasicBlock

	cond types.Value
}

func (w *WaitUntil) Build() types.ScratchBlock {
	return w.BasicBlock.Build("control_wait_until", map[string]types.ScratchInput{
		"CONDITION": w.cond.Build(),
	}, make(map[string]types.ScratchField))
}

type TimerVal struct {
	*BasicBlock
}

func (t *TimerVal) Build() types.ScratchBlock {
	return t.BasicBlock.Build("sensing_timer", make(map[string]types.ScratchInput), make(map[string]types.ScratchField))
}

func (t *TimerVal) ScratchBlockVal() {}

type ResetTimer struct {
	*BasicBlock
}

func (r *ResetTimer) Build() types.ScratchBlock {
	return r.BasicBlock.Build("sensing_resettimer", make(map[string]types.ScratchInput), make(map[string]types.ScratchField))
}

func (b *Blocks) NewWait(time types.Value) *Wait {
	return &Wait{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		time:       time,
	}
}

func (b *Blocks) NewWaitUntil(cond types.Value) *WaitUntil {
	return &WaitUntil{
		BasicBlock: newBasicBlock(types.GetRandomString()),
		cond:       cond,
	}
}

func (b *Blocks) NewTimerVal() *TimerVal {
	return &TimerVal{
		BasicBlock: newBasicBlock(types.GetRandomString()),
	}
}

func (b *Blocks) NewResetTimer() *ResetTimer {
	return &ResetTimer{
		BasicBlock: newBasicBlock(types.GetRandomString()),
	}
}
