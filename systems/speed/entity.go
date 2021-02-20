package speed

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

type speedEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*SpeedComponent
}
