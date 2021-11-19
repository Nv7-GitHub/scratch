package scratch

import (
	"testing"

	"github.com/Nv7-Github/scratch/values"
)

func TestVariables(t *testing.T) {
	Clear()
	addBlankBg()

	variable := Stage.AddVariable("variable", "This is a variable.")
	variableMonitor := AddMonitor(variable, MonitorDefault)
	variableMonitor.X = 5
	variableMonitor.Y = 5

	numVar := Stage.AddVariable("numeric_variable", "0")
	numVarMonitor := AddMonitor(numVar, MonitorDefault)
	numVarMonitor.X = 5
	numVarMonitor.Y = 33

	list := Stage.AddList("list", []interface{}{"This is a list.", "It has multiple values.", "It has initial values.", "This value will be deleted."})
	listMonitor := AddListMonitor(list)
	listMonitor.X = 5
	listMonitor.Y = 61

	stack := Stage.NewWhenFlagClicked()

	// Variable blocks
	stack.Add(Stage.NewSetVariable(variable, values.NewStringValue("This is a set variable block.")))
	stack.Add(Stage.NewChangeVariable(numVar, values.NewIntValue(1)))
	stack.Add(Stage.NewChangeVariableVisibility(variable, false))
	stack.Add(Stage.NewChangeVariableVisibility(variable, true))

	// List Blocks
	stack.Add(Stage.NewAddToList(list, values.NewStringValue("This is an add to list block.")))
	stack.Add(Stage.NewDeleteFromList(list, values.NewIntValue(4)))
	stack.Add(Stage.NewInsertInList(list, values.NewIntValue(2), values.NewStringValue("This is an insert in list block.")))
	stack.Add(Stage.NewReplaceInList(list, values.NewIntValue(3), values.NewStringValue("This is a replace in list block.")))
	stack.Add(Stage.NewChangeListVisibility(list, false))
	stack.Add(Stage.NewChangeListVisibility(list, true))

	// TODO: Variable values
	// TODO: List values
	// Do this by making a "demonstrate block" which has no code but accepts a value as input

	saveProject(t.Fatal, "testdata/Variables.sb3")
}
