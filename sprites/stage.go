package sprites

import (
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

func NewStage() *Stage {
	return &Stage{
		BasicSprite: newBasicSprite("Stage"),
		StageStacks: &blocks.StageStacks{Stacks: blocks.NewStacks()},
		StageBlocks: &blocks.StageBlocks{},

		VideoState:           "on",
		TextToSpeechLanguage: nil,
		Broadcasts:           make([]*types.BroadcastMessage, 0),
	}
}

type Stage struct {
	*BasicSprite
	*blocks.StageStacks
	*blocks.StageBlocks

	Broadcasts []*types.BroadcastMessage

	VideoState           string
	TextToSpeechLanguage *string
}

func (s *Stage) Build() *types.ScratchStage {
	basic := s.BasicSprite.Build()
	basic.IsStage = true

	broadcasts := make(map[string]string)
	for _, broadcast := range s.Broadcasts {
		broadcasts[broadcast.ScratchID()] = broadcast.Name
	}
	basic.Broadcasts = broadcasts

	basic.Blocks = s.Stacks.Build()

	return &types.ScratchStage{
		ScratchTargetBase: basic,

		VideoState:           s.VideoState,
		TextToSpeechLanguage: s.TextToSpeechLanguage,
	}
}

func (s *Stage) AddVariable(name string, initialValue interface{}) *types.Variable {
	v := s.BasicSprite.AddVariable(name, initialValue)
	v.Local = false
	v.SetScratchSpriteName("")

	return v
}

func (s *Stage) AddList(name string, initialValues []interface{}) *types.List {
	v := s.BasicSprite.AddList(name, initialValues)
	v.Local = false
	v.SetScratchSpriteName("")

	return v
}

func (s *Stage) NewBroadcast(name string) *types.BroadcastMessage {
	msg := types.NewScratchBroadcastMessage(name, types.GetRandomString())
	s.Broadcasts = append(s.Broadcasts, msg)
	return msg
}
