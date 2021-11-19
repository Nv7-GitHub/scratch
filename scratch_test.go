package scratch

import (
	"archive/zip"
	"image/png"
	"os"
	"testing"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/values"
)

func saveProject(handle func(args ...interface{}), name string) {
	out, err := os.Create(name)
	if err != nil {
		handle(err)
	}
	defer out.Close()

	zip := zip.NewWriter(out)
	defer zip.Close()

	err = Save(zip)
	if err != nil {
		handle(err)
	}
}

func addBlankBg() {
	Stage.AddCostume(assets.CostumeBlank("background1"))
}

func TestBasic(t *testing.T) {
	Clear()

	addBlankBg()

	saveProject(t.Fatal, "testdata/Basic.sb3")
}

func TestImage(t *testing.T) {
	Clear()

	// Load image
	imgF, err := os.Open("testdata/octocat.png")
	if err != nil {
		t.Fatal(err)
	}
	defer imgF.Close()
	img, err := png.Decode(imgF)
	if err != nil {
		t.Fatal(err)
	}

	costume, err := assets.GetCostumeFromImage("Background", img)
	if err != nil {
		t.Fatal(err)
	}
	Stage.AddCostume(costume)

	saveProject(t.Fatal, "testdata/Image.sb3")
}

func TestSprites(t *testing.T) {
	Clear()

	addBlankBg()

	s := sprites.AddSprite("Sprite")
	s.AddCostume(assets.CostumeScratchCat("cat"))
	s.AddCostume(assets.CostumeScratchCatB("cat_b"))

	saveProject(t.Fatal, "testdata/Sprites.sb3")
}

func TestFunctions(t *testing.T) {
	Clear()

	addBlankBg()

	ret := Stage.AddVariable("return", 0)
	retMonitor := AddMonitor(ret, MonitorDefault)
	retMonitor.X = 5
	retMonitor.Y = 5

	fn := Stage.NewFunction(blocks.NewFunctionParameterLabel("add"), blocks.NewFunctionParameterValue("a", blocks.FunctionParameterString, 0), blocks.NewFunctionParameterLabel("to"), blocks.NewFunctionParameterValue("b", blocks.FunctionParameterString, 0))
	fn.Warp = true

	op := Stage.NewMath(fn.Parameters[0], fn.Parameters[1], blocks.MathOperationAdd)
	fn.Add(op)

	set := Stage.NewSetVariable(ret, values.NewBlockValue(op))
	fn.Add(set)

	/*stack := Stage.NewWhenFlagClicked()
	call, err := Stage.NewFunctionCall(fn, values.NewIntValue(1), values.NewIntValue(2))
	if err != nil {
		t.Fatal(err)
	}
	stack.Add(call)*/

	saveProject(t.Fatal, "testdata/Functions.sb3")
}
