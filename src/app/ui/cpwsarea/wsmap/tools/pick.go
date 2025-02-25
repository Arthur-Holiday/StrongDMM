package tools

import (
	"sdmm/util"
)

// ToolPick can be used to select a hovered object instance.
type ToolPick struct {
	tool
}

func (ToolPick) Name() string {
	return TNPick
}

func newPick() *ToolPick {
	return &ToolPick{}
}

func (ToolPick) AltBehaviour() bool {
	return false
}

func (t ToolPick) onStart(util.Point) {
	if hoveredInstance := ed.HoveredInstance(); hoveredInstance != nil {
		ed.InstanceSelect(hoveredInstance)
	}
}
