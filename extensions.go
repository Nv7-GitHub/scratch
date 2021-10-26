package scratch

type Extension int

const (
	ExtensionPen Extension = iota
	ExtensionWedo2
	ExtensionMusic
	ExtensionMicrobit
	ExtensionText2Speech
	ExtensionTranslate
	ExtensionVideoSensing
	ExtensionEV3
	ExtensionMakeyMakey
	ExtensionBoost
	ExtensionGDXFor
)

var extensionNames = map[Extension]string{
	ExtensionPen:          "pen",
	ExtensionWedo2:        "wedo2",
	ExtensionMusic:        "music",
	ExtensionMicrobit:     "microbit",
	ExtensionText2Speech:  "text2speech",
	ExtensionTranslate:    "translate",
	ExtensionVideoSensing: "videoSensing",
	ExtensionEV3:          "ev3",
	ExtensionMakeyMakey:   "makeymakey",
	ExtensionBoost:        "boost",
	ExtensionGDXFor:       "gdxfor",
}
