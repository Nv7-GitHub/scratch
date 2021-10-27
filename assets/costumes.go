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

var costumes = make([]*Costume, 0)

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
