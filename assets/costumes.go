package assets

import (
	"io"

	"github.com/Nv7-Github/scratch/types"
)

var costumes = make([]*Costume, 0)

type Costume struct {
	Name            string
	RotationCenterX int
	RotationCenterY int

	id     string
	format string
	data   io.Reader
}

func (c *Costume) ScratchID() string {
	return c.id
}

func (c *Costume) ScratchFormat() string {
	return c.format
}

func (c *Costume) Filename() string {
	return c.id + "." + c.format
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
