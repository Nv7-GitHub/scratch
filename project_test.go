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
	say := s.NewSayForTimeBlock("Hello, World!", 0.5)
	stack.Add(say)

	// Variables
	variable := s.AddVariable("variable", "This is a variable. It will be changed once the say completes.")
	m := AddMonitor(variable, MonitorDefault)
	m.X = 5
	m.Y = 5
	set := s.NewSetVariable(variable, values.NewStringValue("This variable has been changed."))
	stack.Add(set)

	// Loop
	loop := s.NewRepeat(10)
	stack.Add(loop)
	say = s.NewSayForTimeBlock("Hi", 0.5)
	loop.Add(say)

	// Global var & compare
	global := Stage.AddVariable("onelessthantwo", "Not calculated yet!")
	m = AddMonitor(global, MonitorDefault)
	m.X = 5
	m.Y = 10
	lt := s.NewCompare(values.NewIntValue(1), values.NewIntValue(2), blocks.CompareLessThan)
	set = s.NewSetVariable(global, values.NewBlockValue(lt))
	stack.Add(set)

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
