package types

type ScratchProject struct {
	Targets    []ScratchTarget  `json:"targets"`
	Monitors   []ScratchMonitor `json:"monitors"`
	Extensions []string         `json:"extensions"`
	Metadata   ScratchMetadata  `json:"meta"`
}

type ScratchMetadata struct {
	SemVer string `json:"semver"`
	VM     string `json:"vm"`
	Agent  string `json:"agent"`
}

type ScratchMonitor struct {
	// Data
	ID         string            `json:"id"`         // ID of variable it shows
	Mode       string            `json:"mode"`       // Default, Large, Slider, List
	Opcode     string            `json:"opcode"`     // Seems to be `data_variable` or `data_listcontents`
	Params     map[string]string `json:"params"`     // Parameters, can have VARIABLE => variable name, LIST => list name
	SpriteName *string           `json:"spriteName"` // Name of sprite its associated to, can be null or string
	Value      interface{}       `json:"value"`      // Value on monitor

	// Characteristics
	Width   int  `json:"width"`
	Height  int  `json:"height"`
	X       int  `json:"x"`
	Y       int  `json:"y"`
	Visible bool `json:"visible"`

	// Slider
	SliderMin  *int `json:"sliderMin,omitempty"`
	SliderMax  *int `json:"sliderMax,omitempty"`
	IsDiscrete bool `json:"isDiscrete"` // Does slider allow integer values?
}
