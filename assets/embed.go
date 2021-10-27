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
	return GetCostume(name, strings.NewReader(scratchCatSvg), CostumeFormatSVG)
}

func CostumeScratchCatB(name string) *Costume {
	return GetCostume(name, strings.NewReader(catBSvg), CostumeFormatSVG)
}
