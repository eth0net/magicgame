package control

import (
	"github.com/EngoEngine/ecs"
)

// ControlFace enforces type safe access to the underlying ControlComponent.
type ControlFace interface {
	GetControlComponent() *ControlComponent
}

// Controlable defines requirements for adding
// entities to the ControlSystem automatically.
type Controlable interface {
	ecs.BasicFace
	ControlFace
}
