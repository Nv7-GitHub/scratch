package assets

import (
	"bytes"
	"image"
	"image/png"
	"io"
	"math"

	"github.com/Nv7-Github/scratch/types"
	"github.com/nfnt/resize"
)

var costumes []*Costume

type Costume struct {
	Name             string
	RotationCenterX  int
	RotationCenterY  int
	BitmapResolution int

	id     string
	format string
	data   io.Reader
}

func (c *Costume) filename() string {
	return c.id + "." + c.format
}

func (c *Costume) Build() types.ScratchCostume {
	return types.ScratchCostume{
		ScratchAsset: types.ScratchAsset{
			AssetID:    c.id,
			Name:       c.Name,
			Md5Ext:     c.filename(),
			DataFormat: c.format,
		},
		BitmapResolution: c.BitmapResolution,
		RotationCenterX:  c.RotationCenterX,
		RotationCenterY:  c.RotationCenterY,
	}
}

type CostumeFormat int

const (
	CostumeFormatSVG CostumeFormat = iota
	CostumeFormatPNG
)

var costumeFormatNames = map[CostumeFormat]string{
	CostumeFormatSVG: "svg",
	CostumeFormatPNG: "png",
}

func GetCostume(name string, data io.Reader, format CostumeFormat) *Costume {
	costume := &Costume{
		Name:   name,
		data:   data,
		format: costumeFormatNames[format],
		id:     types.GetRandomString(),
	}
	costumes = append(costumes, costume)

	return costume
}

func GetCostumeFromImage(name string, img image.Image) (*Costume, error) {
	// Calculate ratios and center
	// The way it works is:
	// 1. Calculate the ratio of the image's X and Y to the scratch resolution (480x360)
	// 2. If the ratio on the X is biggest, then the center X is max X and center Y is (ratio of Y to X) * max Y
	// 3. If the ratio on the Y is biggest, then the center X is (ratio of X to Y) * max X and center Y is max Y
	// Example: If you had a 2048x2048 image, the Y ratio would be larger, and the center would be (360, 360)
	// 4. If the image is over 2x the max scratch resolution, resize it to fit within 2x that resolution and set the DPI (BitmapResolution) to 2
	diffX := math.Ceil(float64(img.Bounds().Dx()) / types.ScratchResolutionX)
	diffY := math.Ceil(float64(img.Bounds().Dy()) / types.ScratchResolutionY)
	diff := 1
	rotX := img.Bounds().Dx() / 2
	rotY := img.Bounds().Dy() / 2
	if diffX > 1 || diffY > 1 {
		if diffX > diffY {
			diff = int(diffX)
			rotX = types.ScratchResolutionX
			rotY = int(float64(img.Bounds().Dy())/float64(img.Bounds().Dx())) * types.ScratchResolutionX
		} else {
			diff = int(diffY)
			rotX = int(float64(img.Bounds().Dx())/float64(img.Bounds().Dy())) * types.ScratchResolutionY
			rotY = types.ScratchResolutionY
		}
	}

	if diff > types.MaxBitmapResolution {
		diff = types.MaxBitmapResolution
		img = resize.Thumbnail(types.ScratchResolutionX*types.MaxBitmapResolution, types.ScratchResolutionY*types.MaxBitmapResolution, img, resize.Lanczos3)
	}

	// Get buf
	buf := bytes.NewBuffer(nil)
	err := png.Encode(buf, img)
	if err != nil {
		return nil, err
	}

	costume := GetCostume(name, buf, CostumeFormatPNG)
	costume.BitmapResolution = diff
	costume.RotationCenterX = rotX
	costume.RotationCenterY = rotY
	return costume, nil
}
