package character

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

// CharacterSystem handles character-specific operations.
type CharacterSystem struct {
	entities []characterEntity
}

// Add an entity to the CharacterSystem.
func (cs *CharacterSystem) Add(
	basic *ecs.BasicEntity,
	anim *common.AnimationComponent,
	space *common.SpaceComponent,
	char *CharacterComponent,
	speed *speed.SpeedComponent,
) {
	if cs.entities == nil {
		cs.entities = []characterEntity{}
	}
	entity := characterEntity{basic, anim, space, char, speed}
	cs.entities = append(cs.entities, entity)
	char.Schedule.Actions = append(char.Schedule.Actions, Action{})
}

// AddByInterface adds entities to the system via Characterable interface.
func (cs *CharacterSystem) AddByInterface(i ecs.Identifier) {
	e := i.(Characterable)
	cs.Add(
		e.GetBasicEntity(),
		e.GetAnimationComponent(),
		e.GetSpaceComponent(),
		e.GetCharacterComponent(),
		e.GetSpeedComponent(),
	)
}

// Remove an entity from the CharacterSystem.
func (cs *CharacterSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range cs.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		cs.entities = append(cs.entities[:del], cs.entities[del+1:]...)
	}
}

// Update the CharacterSystem this frame.
func (cs *CharacterSystem) Update(dt float32) {
	for _, e := range cs.entities {
		e.setAnimation()
	}
}
