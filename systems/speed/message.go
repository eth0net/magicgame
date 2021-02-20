package speed

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

// SpeedMessage updates the speed of an object within the system.
type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

// Type returns a unique string representation for SpeedMessage.
func (SpeedMessage) Type() string {
	return SpeedMessageType
}
