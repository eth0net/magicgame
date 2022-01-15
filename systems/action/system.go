package action

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

// ActionSystem handles action-specific operations.
type ActionSystem struct {
	entities []actionEntity
}

// Add an entity to the ActionSystem.
func (as *ActionSystem) Add(
	basic *ecs.BasicEntity,
	anim *common.AnimationComponent,
	space *common.SpaceComponent,
	action *ActionComponent,
	speed *speed.SpeedComponent,
) {
	if as.entities == nil {
		as.entities = []actionEntity{}
	}
	entity := actionEntity{basic, anim, space, action, speed}
	as.entities = append(as.entities, entity)
}

// AddByInterface adds entities to the system via Actionable interface.
func (as *ActionSystem) AddByInterface(i ecs.Identifier) {
	e := i.(Actionable)
	as.Add(
		e.GetBasicEntity(),
		e.GetAnimationComponent(),
		e.GetSpaceComponent(),
		e.GetActionComponent(),
		e.GetSpeedComponent(),
	)
}

// Remove an entity from the ActionSystem.
func (as *ActionSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range as.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		as.entities = append(as.entities[:del], as.entities[del+1:]...)
	}
}

// Update the ActionSystem this frame.
func (as *ActionSystem) Update(dt float32) {
	for _, e := range as.entities {
		e.runSchedule(dt)
		e.setAnimation()
	}
}
