package sprites

import (
	"github.com/Nv7-Github/scratch/assets"
	"github.com/Nv7-Github/scratch/blocks"
	"github.com/Nv7-Github/scratch/types"
)

func newBasicSprite(name string) *BasicSprite {
	return &BasicSprite{
		Name:      name,
		Variables: make(map[string]*types.Variable),
		Lists:     make(map[string]*types.List),
		Costumes:  make([]*assets.Costume, 0),
		Sounds:    make([]*assets.Sound, 0),
		comments:  make(map[string]*Comment),

		Volume: 100,
	}
}

type BasicSprite struct {
	Name string

	Variables map[string]*types.Variable // map[name]*Variable
	Lists     map[string]*types.List     // map[name]*List

	Costume  int
	Costumes []*assets.Costume

	Sounds []*assets.Sound
	Volume int

	comments map[string]*Comment
}

func (b *BasicSprite) AddCostume(costume *assets.Costume) {
	b.Costumes = append(b.Costumes, costume)
}

func (b *BasicSprite) AddSound(sound *assets.Sound) {
	b.Sounds = append(b.Sounds, sound)
}

func (b *BasicSprite) SetComment(block blocks.Block, comment string) *Comment {
	cmt := &Comment{
		id:      types.GetRandomString(),
		blockid: block.ScratchID(),

		Text:      comment,
		Minimized: true,
	}
	b.comments[cmt.id] = cmt
	block.SetCommentID(cmt.id)
	return cmt
}

func (b *BasicSprite) GetComment(block blocks.Block) *Comment {
	return b.comments[*block.CommentID()]
}

type Comment struct {
	id      string
	blockid string

	Text          string
	X, Y          int
	Width, Height int
	Minimized     bool
}

func (b *BasicSprite) AddVariable(name string, initialValue interface{}) *types.Variable {
	variable := &types.Variable{
		Name:         name,
		InitialValue: initialValue,
		Local:        true,
	}
	variable.SetScratchID(types.GetRandomString())
	variable.SetScratchSpriteName(b.Name)
	b.Variables[name] = variable

	return variable
}

func (b *BasicSprite) AddList(name string, initialValues []interface{}) *types.List {
	list := &types.List{
		Name:          name,
		InitialValues: initialValues,
		Local:         true,
	}
	list.SetScratchListID(types.GetRandomString())
	list.SetScratchSpriteName(b.Name)
	b.Lists[name] = list

	return list
}
