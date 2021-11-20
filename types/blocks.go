package types

type ScratchBlock struct {
	Opcode   string                  `json:"opcode"`
	Next     *string                 `json:"next"`   // string or null, id of next
	Parent   *string                 `json:"parent"` // string or null - from scratch wiki: If the block is a stack block and is preceded, this is the ID of the preceding block. If the block is the first stack block in a C mouth, this is the ID of the C block. If the block is an input to another block, this is the ID of that other block. Otherwise it is null.
	Inputs   map[string]ScratchInput `json:"inputs"` // map[inputName]InputValue, use Shadow for constand and Obscured for block in it
	Fields   map[string]ScratchField `json:"fields"` // map[fieldName]FieldValue
	Shadow   bool                    `json:"shadow"`
	TopLevel bool                    `json:"topLevel"` // False if the block has a parent and true otherwise.
	Mutation *ScratchMutation        `json:"mutation,omitempty"`

	X       *int    `json:"x,omitempty"` // If toplevel
	Y       *int    `json:"y,omitempty"`
	Comment *string `json:"comment,omitempty"` // ID of comment if it has one
}

type ScratchMutation struct {
	TagName          string  `json:"tagName"`          // always "mutation"
	Children         []bool  `json:"children"`         // seems to be just an empty array
	ProcCode         string  `json:"proccode"`         // name of block, has parameters in it like "add %s %s label %b" where %s is string/number and %b is boolean
	ArgumentIDs      string  `json:"argumentids"`      // []string, but marshal with JSON (use MarshalStringArray)
	ArgumentNames    *string `json:"argumentnames"`    // []string, but marshal with JSON (use MarshalStringArray)
	ArgumentDefaults string  `json:"argumentdefaults"` // []interface{}, but marshal each element with JSON into []string and then marshal that with JSON (use MarshalInterfaceArray)
	Warp             string  `json:"warp"`             // run without screen refresh? (string of bool, with JSON)
}

type ScratchInput []interface{}

func NewScratchInputShadow(val ScratchValue) ScratchInput {
	return ScratchInput{1, val}
}

func NewScratchInputNoShadow(val string) ScratchInput {
	return ScratchInput{2, val}
}

func NewScratchInputNoShadowBlock(val ScratchValue) ScratchInput {
	return ScratchInput{2, val}
}

func NewScratchInputObscured(val ScratchValue, below ScratchValue) ScratchInput {
	return ScratchInput{3, val, below}
}

func NewScratchInputStack(firstId string) ScratchInput {
	return ScratchInput{2, firstId}
}

type ScratchValue interface {
	mustBeInput()
}

type scratchInputBlock string

func (s *scratchInputBlock) mustBeInput() {}

type scratchInputVal []interface{}

func NewScratchBlockInput(id string) ScratchValue {
	blk := scratchInputBlock(id)
	return &blk
}

func (s *scratchInputVal) mustBeInput() {}

func NewScratchFloat(val float64) ScratchValue {
	return &scratchInputVal{4, val}
}

func NewScratchPosFloat(val float64) ScratchValue {
	return &scratchInputVal{5, val}
}

func NewScratchPosInt(val int) ScratchValue {
	return &scratchInputVal{6, val}
}

func NewScratchInt(val int) ScratchValue {
	return &scratchInputVal{7, val}
}

func NewScratchAngle(val int) ScratchValue {
	return &scratchInputVal{8, val}
}

// NOTE: Requires with #
func NewScratchColor(hex string) ScratchValue {
	return &scratchInputVal{9, hex}
}

func NewScratchString(val string) ScratchValue {
	return &scratchInputVal{10, val}
}

func NewScratchBroadcast(name, id string) ScratchValue {
	return &scratchInputVal{11, name, id}
}

func NewScratchVariable(name, id string) ScratchValue {
	return &scratchInputVal{12, name, id}
}

func NewScratchList(name, id string) ScratchValue {
	return &scratchInputVal{13, name, id}
}

type ScratchField interface {
	mustBeInputField()
}

type scratchInputField []interface{}

func (s *scratchInputField) mustBeInputField() {}

func NewScratchValueFieldVariable(name, id string) ScratchField {
	return &scratchInputField{name, id}
}

func NewScratchFieldBroadcast(name, id string) ScratchField {
	return &scratchInputField{name, id}
}

func NewScratchFieldParamName(name string) ScratchField {
	return &scratchInputField{name, nil}
}
