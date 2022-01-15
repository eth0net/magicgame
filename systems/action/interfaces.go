package action

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

// ActionFace enforces type safe access to underlying ActionComponent.
type ActionFace interface {
	GetActionComponent() *ActionComponent
}

// Actionable defines requirements for adding
// entities to the ActionSystem automatically.
type Actionable interface {
	ecs.BasicFace
	common.AnimationFace
	common.SpaceFace
	ActionFace
	speed.SpeedFace
}
