package scratch

import (
	"testing"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/values"
)

func createProject(handler func(args ...interface{})) {
	Clear()

	addBlankBg()

	s := sprites.AddSprite("Sprite")
	s.AddCostume(assets.CostumeScratchCat("cat"))

	stack := s.NewWhenFlagClicked()
	say := s.NewSayForTimeBlock("Hello, World!", 0.5)
	stack.Add(say)

	variable := s.AddVariable("variable", "This is a variable. It will be changed once the say completes.")
	AddMonitor(variable, MonitorDefault)
	set := s.NewSetVariable(variable, values.NewStringValue("This variable has been changed."))
	stack.Add(set)

	loop := s.NewRepeat(10)
	stack.Add(loop)

	say = s.NewSayForTimeBlock("Hi", 0.5)
	loop.Add(say)

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
