package speed

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

// SpeedFace enforces type safe access to the underlying SpeedComponent.
type SpeedFace interface {
	GetSpeedComponent() *SpeedComponent
}

// Speedable defines requirements for adding
// entities to the SpeedSystem automatically.
type Speedable interface {
	ecs.BasicFace
	common.AnimationFace
	common.SpaceFace
	SpeedFace
}
