package scratch

import (
	"archive/zip"
	"image/png"
	"os"
	"testing"

	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/sprites"
)

func saveProject(t *testing.T, name string) {
	out, err := os.Create(name)
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	zip := zip.NewWriter(out)
	defer zip.Close()

	err = Save(zip)
	if err != nil {
		t.Fatal(err)
	}
}

func addBlankBg() {
	Stage.AddCostume(assets.CostumeBlank("background1"))
}

func TestBasic(t *testing.T) {
	Clear()

	addBlankBg()

	saveProject(t, "testdata/Basic.sb3")
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

	saveProject(t, "testdata/Image.sb3")
}

func TestVariables(t *testing.T) {
	Clear()
	addBlankBg()

	variable := Stage.AddVariable("variable", "This is a variable.")
	variableMonitor := AddMonitor(variable, MonitorDefault)
	variableMonitor.X = 5
	variableMonitor.Y = 5

	list := Stage.AddList("list", []interface{}{"This is a list.", "It has multiple values.", "It has initial values."})
	listMonitor := AddListMonitor(list)
	listMonitor.X = 5
	listMonitor.Y = 33

	saveProject(t, "testdata/Variables.sb3")
}

func TestSprites(t *testing.T) {
	Clear()

	addBlankBg()

	s := sprites.AddSprite("Sprite")
	s.AddCostume(assets.CostumeScratchCat("cat"))
	s.AddCostume(assets.CostumeScratchCatB("cat_b"))

	saveProject(t, "testdata/Sprites.sb3")
}
