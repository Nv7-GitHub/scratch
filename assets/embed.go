package assets

import (
	_ "embed"
	"strings"
)

//go:embed blank.svg
var blankSvg string

//go:embed cat.svg
var scratchCatSvg string

//go:embed cat_b.svg
var catBSvg string

func CostumeBlank(name string) *Costume {
	return GetCostume(name, strings.NewReader(blankSvg), CostumeFormatSVG)
}

func CostumeScratchCat(name string) *Costume {
	costume := GetCostume(name, strings.NewReader(scratchCatSvg), CostumeFormatSVG)
	costume.RotationCenterX = 48
	costume.RotationCenterY = 50
	return costume
}

func CostumeScratchCatB(name string) *Costume {
	costume := GetCostume(name, strings.NewReader(catBSvg), CostumeFormatSVG)
	costume.RotationCenterX = 46
	costume.RotationCenterY = 53
	return costume
}
