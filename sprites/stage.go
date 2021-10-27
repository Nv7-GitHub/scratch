package sprites

import "github.com/Nv7-Github/scratch/types"

func NewStage() *Stage {
	return &Stage{
		BasicSprite: newBasicSprite("Stage"),

		VideoState:           "on",
		TextToSpeechLanguage: nil,
		Broadcasts:           make([]string, 0),
		broadcastIDs:         make([]string, 0),
	}
}

type Stage struct {
	*BasicSprite

	Broadcasts   []string
	broadcastIDs []string

	VideoState           string
	TextToSpeechLanguage *string
}

func (s *Stage) Build() *types.ScratchStage {
	basic := s.BasicSprite.Build()
	basic.IsStage = true

	broadcasts := make(map[string]string)
	for i, broadcast := range s.Broadcasts {
		broadcasts[s.broadcastIDs[i]] = broadcast
	}
	basic.Broadcasts = broadcasts

	basic.Blocks = make(map[string]types.ScratchBlock) // Empty because no support for blocks yet

	return &types.ScratchStage{
		ScratchTargetBase: basic,

		VideoState:           s.VideoState,
		TextToSpeechLanguage: s.TextToSpeechLanguage,
	}
}

func (s *Stage) AddVariable(name string, initialValue interface{}) *Variable {
	v := s.BasicSprite.AddVariable(name, initialValue)
	(*v).Local = false
	(*v).spriteName = ""

	return v
}

func (s *Stage) AddList(name string, initialValues []interface{}) *List {
	v := s.BasicSprite.AddList(name, initialValues)
	(*v).Local = false
	(*v).spriteName = ""

	return v
}
