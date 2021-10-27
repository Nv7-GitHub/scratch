package scratch

import (
	"archive/zip"
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
