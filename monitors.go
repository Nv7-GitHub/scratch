package scratch

import (
	"github.com/Nv7-Github/scratch/sprites"
	"github.com/Nv7-Github/scratch/types"
)

var Monitors []Monitor

type MonitorMode int

const (
	MonitorDefault MonitorMode = iota
	MonitorLarge
	MonitorSlider
)

var monitorModeNames = map[MonitorMode]string{
	MonitorDefault: "default",
	MonitorLarge:   "large",
	MonitorSlider:  "slider",
}

func AddMonitor(v *sprites.Variable, mode MonitorMode) *VariableMonitor {
	m := &VariableMonitor{
		Variable:  v,
		Mode:      mode,
		SliderMax: 100,
		visible:   true,
	}
	Monitors = append(Monitors, m)
	return m
}

func AddListMonitor(v *sprites.List) *ListMonitor {
	m := &ListMonitor{
		List:    v,
		visible: true,
	}
	Monitors = append(Monitors, m)
	return m
}

type Monitor interface {
	Build() types.ScratchMonitor
	Show()
	Hide()
}

type VariableMonitor struct {
	Variable             *sprites.Variable
	X, Y                 int
	Width, Height        int
	Mode                 MonitorMode
	SliderMin, SliderMax int

	visible bool
}

func (v *VariableMonitor) Show() {
	v.visible = true
}

func (v *VariableMonitor) Hide() {
	v.visible = false
}

type ListMonitor struct {
	List          *sprites.List
	X, Y          int
	Width, Height int

	visible bool
}

func (l *ListMonitor) Show() {
	l.visible = true
}

func (l *ListMonitor) Hide() {
	l.visible = false
}

func (v *VariableMonitor) Build() types.ScratchMonitor {
	m := types.ScratchMonitor{
		ID:     v.Variable.ScratchID(),
		Mode:   monitorModeNames[v.Mode],
		Opcode: "data_variable",
		Params: map[string]string{"VARIABLE": v.Variable.Name},
		Value:  v.Variable.InitialValue,

		X:       v.X,
		Y:       v.Y,
		Width:   v.Width,
		Height:  v.Height,
		Visible: v.visible,

		SliderMin:  &v.SliderMin,
		SliderMax:  &v.SliderMax,
		IsDiscrete: true,
	}
	if v.Variable.Local {
		name := v.Variable.SpriteName()
		m.SpriteName = &name
	}
	return m
}

func (l *ListMonitor) Build() types.ScratchMonitor {
	m := types.ScratchMonitor{
		ID:     l.List.ScratchID(),
		Mode:   "list",
		Opcode: "data_listcontents",
		Params: map[string]string{"LIST": l.List.Name},
		Value:  l.List.InitialValues,

		X:       l.X,
		Y:       l.Y,
		Width:   l.Width,
		Height:  l.Height,
		Visible: l.visible,

		IsDiscrete: false,
	}
	if l.List.Local {
		name := l.List.SpriteName()
		m.SpriteName = &name
	}
	return m
}
