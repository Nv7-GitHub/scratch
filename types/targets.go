package types

type ScratchTarget struct {
	IsStage        bool                            `json:"isStage"`
	Name           string                          `json:"name"`
	Variables      map[string]ScratchVariableValue `json:"variables"`
	Lists          map[string]ScratchVariableValue `json:"lists"`
	Broadcasts     map[string]string               `json:"broadcasts"` // map[id]name, is an empty map on everything but stage
	Blocks         map[string]ScratchBlock         `json:"blocks"`     // map[id]block
	Comments       map[string]ScratchComment       `json:"comments"`   // map[id]comment
	CurrentCostume int                             `json:"currentCostume"`
	Costumes       []ScratchCostume                `json:"costumes"`
	Sounds         []ScratchSound                  `json:"sounds"`
	Volume         int                             `json:"volume"`
}

type ScratchStage struct {
	ScratchTarget

	TextToSpeechLanguage *string `json:"textToSpeechLanguage"` // string or null
	VideoState           string  `json:"videoState"`           // seems to be "on" always
}

type ScratchSprite struct {
	ScratchTarget

	LayerOrder    int    `json:"layerOrder"` // Layer on screen
	X             int    `json:"x"`
	Y             int    `json:"y"`
	Size          int    `json:"size"`      // 0-100 percent, can go past 100 to scale up
	Direction     int    `json:"direction"` // degrees
	Draggable     bool   `json:"draggable"`
	RotationStyle string `json:"rotationStyle"` // Scratch rotation style, can be `all around`, `left-right`, and `don't rotate`
}

type ScratchVariableValue []interface{} // name (string), value (int, string, list, etc.)

type ScratchAsset struct {
	AssetID    string `json:"assetId"` // Name of file without extension
	Name       string `json:"name"`
	Md5Ext     string `json:"md5ext"`     // AssetID + extension
	DataFormat string `json:"dataFormat"` // `svg` if costume, `wav` if sound, `png` if image
}

type ScratchCostume struct {
	ScratchAsset

	BitmapResolution int `json:"bitmapResolution"` // Usually 1
	RotationCenterX  int `json:"rotationCenterX"`
	RotationCenterY  int `json:"rotationCenterY"`
}

type ScratchSound struct {
	ScratchAsset

	Format      string `json:"format"`      // Not sure what this is, seems to be an empty string
	Rate        int    `json:"rate"`        // Sampling rate in hertz
	SampleCount int    `json:"sampleCount"` // Sample count of WAV
}

type ScratchComment struct {
	BlockID   string `json:"blockId"`
	X         int    `json:"x"`
	Y         int    `json:"y"`
	Width     int    `json:"width"`
	Height    int    `json:"height"`
	Minimized bool   `json:"minimized"`
	Text      string `json:"text"`
}
