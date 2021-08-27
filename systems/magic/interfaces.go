package magic

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

// Face enforces type safe access to the underlying Component.
type Face interface {
	GetComponent() *Component
}

// Able defines requirements for adding entities to the System automatically.
type Able interface {
	ecs.BasicFace
	Face
	common.SpaceFace
}
