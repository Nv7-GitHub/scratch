package sprites

import "github.com/Nv7-Github/scratch/types"

func NewStage() *Stage {
	return &Stage{
		BasicSprite: newBasicSprite(),

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
