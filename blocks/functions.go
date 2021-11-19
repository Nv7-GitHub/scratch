package blocks

import (
	"errors"
	"strings"

	"github.com/Nv7-Github/scratch/types"
)

// Types
type FunctionParameterType int

const (
	FunctionParameterString FunctionParameterType = iota
	FunctionParameterBool
)

type FunctionParameter interface {
	ProcCode() string
}

type FunctionParameterValue interface {
	FunctionParameter

	Name() string
	Type() FunctionParameterType
	Default() interface{}
}

type ParamType struct {
	name         string
	id           string
	keyId        string
	typ          FunctionParameterType
	initialValue interface{}
}

type ScratchParamValue struct {
	id string
}

func (s *ScratchParamValue) Build() types.ScratchInput {
	return types.NewScratchInputShadow(types.NewScratchBlockInput(s.id))
}

// Code
type Function struct {
	*BasicStack

	X, Y int
	Warp bool // Run without screen refresh

	Parameters []types.Value
	paramTypes []ParamType
	procCode   string
}

func (f *Function) Build() map[string]types.ScratchBlock {
	customBlkId := types.GetRandomString()
	id := types.GetRandomString()
	blk := types.ScratchBlock{
		Opcode:   "procedures_definition",
		Inputs:   map[string]types.ScratchInput{"custom_block": types.NewScratchInputShadow(types.NewScratchBlockInput(customBlkId))},
		Fields:   make(map[string]types.ScratchField),
		Shadow:   false,
		TopLevel: true,
		X:        &f.X,
		Y:        &f.Y,
	}
	stack := f.BasicStack.Build(blk, id)

	// Create definition
	argIDs := make([]string, len(f.paramTypes))
	inps := make(map[string]types.ScratchInput)
	for i, v := range f.paramTypes {
		argIDs[i] = v.keyId
		inps[v.keyId] = types.NewScratchInputShadow(types.NewScratchBlockInput(v.id))
	}

	// Get other data
	argNames := make([]string, len(f.paramTypes))
	argDefaults := make([]interface{}, len(f.paramTypes))
	for i, v := range f.paramTypes {
		argNames[i] = v.name
		argDefaults[i] = v.initialValue
	}

	argNameVal := types.MarshalStringArray(argNames)
	stack[customBlkId] = types.ScratchBlock{
		Opcode:   "procedures_prototype",
		Next:     nil,
		Parent:   &id,
		Inputs:   inps,
		Fields:   make(map[string]types.ScratchField),
		Shadow:   true,
		TopLevel: false,
		Mutation: &types.ScratchMutation{
			TagName:          "mutation",
			Children:         make([]bool, 0),
			ProcCode:         f.procCode,
			ArgumentIDs:      types.MarshalStringArray(argIDs),
			ArgumentNames:    &argNameVal,
			ArgumentDefaults: types.MarshalInterfaceArray(argDefaults),
			Warp:             f.Warp,
		},
	}

	// Add param getters
	for i, parTyp := range f.paramTypes {
		opcode := "argument_reporter_string_number"
		if f.paramTypes[i].typ == FunctionParameterBool {
			opcode = "argument_reporter_boolean"
		}
		stack[parTyp.id] = types.ScratchBlock{
			Opcode: opcode,
			Next:   nil,
			Parent: &customBlkId,
			Inputs: make(map[string]types.ScratchInput),
			Fields: map[string]types.ScratchField{
				"VALUE": types.NewScratchFieldParamName(parTyp.name),
			},
			Shadow:   true,
			TopLevel: false,
		}
	}

	return stack
}

func (s *Stacks) NewFunction(params ...FunctionParameter) *Function {
	procCodePars := make([]string, len(params))
	parCount := 0
	for i, v := range params {
		procCodePars[i] = v.ProcCode()
		_, ok := v.(FunctionParameterValue)
		if ok {
			parCount++
		}
	}
	procCode := strings.Join(procCodePars, " ")

	// Calculate params
	i := 0
	pars := make([]types.Value, parCount)
	parTypes := make([]ParamType, parCount)
	for _, v := range params {
		val, ok := v.(FunctionParameterValue)
		if ok {
			id := types.GetRandomString()
			keyId := types.GetRandomString()
			parTypes[i] = ParamType{
				name:         val.Name(),
				id:           id,
				keyId:        keyId,
				typ:          val.Type(),
				initialValue: val.Default(),
			}
			pars[i] = &ScratchParamValue{
				id: id,
			}
			i++
		}
	}

	return s.addStack(&Function{
		BasicStack: newBasicStack(),

		Parameters: pars,
		paramTypes: parTypes,
		procCode:   procCode,
	}).(*Function)
}

// Function Calls
type FunctionCall struct {
	*BasicBlock

	function *Function
	params   []types.Value
}

func (f *FunctionCall) Build() types.ScratchBlock {
	inpMap := make(map[string]types.ScratchInput, len(f.params))
	argumentIds := make([]string, len(f.function.paramTypes))
	for i, param := range f.params {
		inpMap[f.function.paramTypes[i].keyId] = param.Build()
		argumentIds[i] = f.function.paramTypes[i].keyId
	}

	mutation := types.ScratchMutation{
		TagName:     "mutation",
		Children:    make([]bool, 0),
		ProcCode:    f.function.procCode,
		ArgumentIDs: types.MarshalStringArray(argumentIds),
		Warp:        f.function.Warp,
	}
	blk := f.BasicBlock.Build("procedures_call", inpMap, make(map[string]types.ScratchField))
	blk.Mutation = &mutation
	return blk
}

func (b *Blocks) NewFunctionCall(fn *Function, params ...types.Value) (*FunctionCall, error) {
	if len(params) != len(fn.paramTypes) {
		return nil, errors.New("incorrect number of arguments")
	}
	return &FunctionCall{
		BasicBlock: newBasicBlock(types.GetRandomString()),

		function: fn,
	}, nil
}

// Function parameter types
type FunctionParameterLabel struct {
	Label string
}

func (f *FunctionParameterLabel) ProcCode() string {
	return f.Label
}

type FunctionParameterInput struct {
	name         string
	typ          FunctionParameterType
	initialValue interface{}
}

func (f *FunctionParameterInput) ProcCode() string {
	if f.typ == FunctionParameterString {
		return "%s"
	}
	return "%b"
}

func (f *FunctionParameterInput) Name() string {
	return f.name
}

func (f *FunctionParameterInput) Type() FunctionParameterType {
	return f.typ
}

func (f *FunctionParameterInput) Default() interface{} {
	return f.initialValue
}

func NewFunctionParameterLabel(text string) *FunctionParameterLabel {
	return &FunctionParameterLabel{
		Label: text,
	}
}

func NewFunctionParameterValue(name string, typ FunctionParameterType, initialValue interface{}) *FunctionParameterInput {
	return &FunctionParameterInput{
		name:         name,
		typ:          typ,
		initialValue: initialValue,
	}
}
