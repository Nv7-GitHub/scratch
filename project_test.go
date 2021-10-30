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

	// Variable
	iter := s.AddVariable("i", 0)
	AddMonitor(iter, MonitorDefault)

	// FOR loop (i := 0; i < 10; i++)
	initialize := s.NewSetVariable(iter, values.NewIntValue(0)) // i := 0
	stack.Add(initialize)

	var condition blocks.Block = s.NewCompare(values.NewVariableValue(iter), values.NewIntValue(10), blocks.CompareLessThan) // i < 10
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
