package scratch

import (
	"archive/zip"
	"image/png"
	"os"
	"testing"

	"github.com/Nv7-Github/scratch/assets"
)

func TestBasic(t *testing.T) {
	out, err := os.Create("Project.sb3")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

	Stage.AddCostume(assets.CostumeBlank("background1"))

	zip := zip.NewWriter(out)
	defer zip.Close()

	err = Save(zip)
	if err != nil {
		t.Fatal(err)
	}
}

func TestImage(t *testing.T) {
	out, err := os.Create("Project.sb3")
	if err != nil {
		t.Fatal(err)
	}
	defer out.Close()

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

	zip := zip.NewWriter(out)
	defer zip.Close()

	err = Save(zip)
	if err != nil {
		t.Fatal(err)
	}
}
