package types

type ScratchBlock struct {
	Opcode   string                  `json:"opcode"`
	Next     *string                 `json:"next"`   // string or null, id of next
	Parent   *string                 `json:"parent"` // string or null - from scratch wiki: If the block is a stack block and is preceded, this is the ID of the preceding block. If the block is the first stack block in a C mouth, this is the ID of the C block. If the block is an input to another block, this is the ID of that other block. Otherwise it is null.
	Inputs   map[string]ScratchInput `json:"inputs"` // map[inputName]InputValue
	Fields   map[string][]string     `json:"fields"` // Not sure what this is, empty from what I've seen
	Shadow   bool                    `json:"shadow"`
	TopLevel bool                    `json:"topLevel"` // False if the block has a parent and true otherwise.

	X       *int    `json:"x,omitempty"` // If toplevel
	Y       *int    `json:"y,omitempty"`
	Comment *string `json:"comment,omitempty"` // ID of comment if it has one
}

type ScratchInput interface {
	mustBeInput()
}

type scratchInputVal []interface{}

func (s *scratchInputVal) mustBeInput() {}

func NewScratchFloat(val float64) ScratchInput {
	return &scratchInputVal{4, val}
}

func NewScratchPosFloat(val float64) ScratchInput {
	return &scratchInputVal{5, val}
}

func NewScratchPosInt(val int) ScratchInput {
	return &scratchInputVal{6, val}
}

func NewScratchInt(val int) ScratchInput {
	return &scratchInputVal{7, val}
}

func NewScratchAngle(val int) ScratchInput {
	return &scratchInputVal{8, val}
}

// NOTE: Requires with #
func NewScratchColor(hex string) ScratchInput {
	return &scratchInputVal{9, hex}
}

func NewScratchString(val string) ScratchInput {
	return &scratchInputVal{10, val}
}

func NewScratchBroadcast(name, id string) ScratchInput {
	return &scratchInputVal{11, name, id}
}

func NewScratchVariable(name, id string) ScratchInput {
	return &scratchInputVal{12, name, id}
}
