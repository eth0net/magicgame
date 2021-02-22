package control

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/raziel2244/magicgame/systems/speed"
)

// ControlSystem handles player input.
type ControlSystem struct {
	entities []controlEntity
}

// New initialises SpeedSystem when it's added to the world.
func (ss *ControlSystem) New(*ecs.World) {
	engo.Input.RegisterButton(ButtonUp, engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton(ButtonDown, engo.KeyS, engo.KeyArrowDown)
	engo.Input.RegisterButton(ButtonLeft, engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton(ButtonRight, engo.KeyD, engo.KeyArrowRight)
	engo.Input.RegisterButton(ButtonSprint, engo.KeyLeftShift)

	engo.Input.RegisterAxis(
		AxisVertical,
		engo.AxisKeyPair{Min: engo.KeyW, Max: engo.KeyS},
		engo.AxisKeyPair{Min: engo.KeyArrowUp, Max: engo.KeyArrowDown},
	)

	engo.Input.RegisterAxis(
		AxisHorizontal,
		engo.AxisKeyPair{Min: engo.KeyA, Max: engo.KeyD},
		engo.AxisKeyPair{Min: engo.KeyArrowLeft, Max: engo.KeyArrowRight},
	)
}

// Add an entity to the ControlSystem.
func (cs *ControlSystem) Add(
	b *ecs.BasicEntity,
	c *ControlComponent,
) {
	if cs.entities == nil {
		cs.entities = []controlEntity{}
	}
	cs.entities = append(cs.entities, controlEntity{b, c})
}

// AddByInterface adds entities to the system via Controlable interface.
func (cs *ControlSystem) AddByInterface(i ecs.Identifier) {
	e := i.(Controlable)
	cs.Add(
		e.GetBasicEntity(),
		e.GetControlComponent(),
	)
}

// Remove an entity from the ControlSystem.
func (cs *ControlSystem) Remove(b ecs.BasicEntity) {
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

// Update the ControlSystem this frame.
func (cs *ControlSystem) Update(dt float32) {
	for _, e := range cs.entities {
		if !e.ControlComponent.Enabled {
			continue
		}
		if vector, changed := e.speed(); changed {
			vector, _ = vector.Normalize()
			vector.MultiplyScalar(dt)

			if engo.Input.Button(ButtonSprint).Down() {
				vector.MultiplyScalar(2)
			}

			engo.Mailbox.Dispatch(speed.SpeedMessage{e.BasicEntity, vector})
		}
	}
}
