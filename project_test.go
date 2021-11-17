package scratch

import (
	"testing"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/values"
)

func createProject(handler func(args ...interface{})) {
	// Setup
	Clear()
	addBlankBg()
	s := sprites.AddSprite("Sprite")
	s.AddCostume(assets.CostumeScratchCat("cat"))

	// Stack and hello world
	stack := s.NewWhenFlagClicked()
	say := s.NewSayForTimeBlock(values.NewStringValue("Hello, World!"), 0.5)
	stack.Add(say)

	s.SetComment(say, "Hello World Example")

	// Variable
	iter := s.AddVariable("i", 0)
	mon := AddMonitor(iter, MonitorDefault)
	mon.X = 5
	mon.Y = 5

	// FOR loop (i := 0; i < 10; i++)
	initialize := s.NewSetVariable(iter, values.NewIntValue(0)) // i := 0
	stack.Add(initialize)
	s.SetComment(initialize, "Loop example")

	var condition blocks.BlockVal = s.NewCompare(values.NewVariableValue(iter), values.NewIntValue(10), blocks.CompareLessThan) // i < 10
	stack.Add(condition)
	condition = s.NewNot(values.NewBlockValue(condition)) // ! since repeat until
	stack.Add(condition)

	loop := s.NewRepeatUntil(values.NewBlockValue(condition))
	stack.Add(loop)

	// Add contents of loop
	say = s.NewSayForTimeBlock(values.NewVariableValue(iter), 0.25) // say(i)
	loop.Add(say)

	incr := s.NewSetVariable(iter, values.NewIntValue(1)) // i++
	incr.Change = true
	loop.Add(incr)

	// Concurrency
	mainThreadVar := s.AddVariable("mainThreadCounter", 0)
	concurrentVar := s.AddVariable("concurrentCounter", 0)
	msg := Stage.NewBroadcast("code")

	mon = AddMonitor(mainThreadVar, MonitorDefault)
	mon.X = 5
	mon.Y = 33
	mon = AddMonitor(concurrentVar, MonitorDefault)
	mon.X = 5
	mon.Y = 61

	broadcast := s.NewBroadcast(msg)
	stack.Add(broadcast)
	s.SetComment(broadcast, "Concurrency example")

	concurrent := s.NewBroadcastReceive(msg)

	// Conditions for currency
	condMainNot := s.NewCompare(values.NewVariableValue(mainThreadVar), values.NewIntValue(10), blocks.CompareLessThan)
	stack.Add(condMainNot)

	condConcurrentNot := s.NewCompare(values.NewVariableValue(concurrentVar), values.NewIntValue(10), blocks.CompareLessThan)
	concurrent.Add(condConcurrentNot)

	condMain := s.NewNot(values.NewBlockValue(condMainNot))
	stack.Add(condMain)

	condConcurrent := s.NewNot(values.NewBlockValue(condConcurrentNot))
	concurrent.Add(condConcurrent)

	// Loops for concurrency
	mainLoop := s.NewRepeatUntil(values.NewBlockValue(condMain))
	stack.Add(mainLoop)

	concurrentLoop := s.NewRepeatUntil(values.NewBlockValue(condConcurrent))
	concurrent.Add(concurrentLoop)

	mainLoopBlk := s.NewSetVariable(mainThreadVar, values.NewIntValue(1))
	mainLoopBlk.Change = true
	mainLoop.Add(mainLoopBlk)

	concurrentLoopBlk := s.NewSetVariable(concurrentVar, values.NewIntValue(1))
	concurrentLoopBlk.Change = true
	concurrentLoop.Add(concurrentLoopBlk)

	waitMain := s.NewWait(values.NewFloatValue(0.25))
	mainLoop.Add(waitMain)
	waitConcurrent := s.NewWait(values.NewFloatValue(0.25))
	concurrentLoop.Add(waitConcurrent)

	saveProject(handler, "testdata/Project.sb3")
}
func TestProject(t *testing.T) {
	createProject(t.Fatal)
}

func BenchProject(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createProject(b.Fatal)
	}
}
