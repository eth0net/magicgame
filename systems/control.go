package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

var (
	upButton    string = "up"
	downButton  string = "down"
	leftButton  string = "left"
	rightButton string = "right"
)

// ControlComponent stores control input for player entity.
type ControlComponent struct {
	Vertical, Horizontal string
}

// GetControlComponent provides type safe access to ControlComponent.
func (c *ControlComponent) GetControlComponent() *ControlComponent {
	return c
}

type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
}

func (ce *controlEntity) speed() (p engo.Point, changed bool) {
	var (
		upReleased    = engo.Input.Button(upButton).JustReleased()
		downReleased  = engo.Input.Button(downButton).JustReleased()
		leftReleased  = engo.Input.Button(leftButton).JustReleased()
		rightReleased = engo.Input.Button(rightButton).JustReleased()
		upHeld        = engo.Input.Button(upButton).Down()
		downHeld      = engo.Input.Button(downButton).Down()
		leftHeld      = engo.Input.Button(leftButton).Down()
		rightHeld     = engo.Input.Button(rightButton).Down()
	)

	oldX := engo.Input.Axis(ce.Horizontal).Value()
	oldY := engo.Input.Axis(ce.Vertical).Value()

	p.X, p.Y = oldX, oldY

	switch {
	case upReleased, downReleased:
		changed = true
		p.Y = 0
	case leftReleased, rightReleased:
		changed = true
		p.X = 0
	case upHeld:
		p.Y--
	case downHeld:
		p.Y++
	case leftHeld:
		p.X--
	case rightHeld:
		p.X++
	}

	if p.X != oldX || p.Y != oldY {
		changed = true
	}

	if p.X > 1 {
		p.X = 1
	} else if p.X < -1 {
		p.X = -1
	}
	if p.Y > 1 {
		p.Y = 1
	} else if p.Y < -1 {
		p.Y = -1
	}

	return p, changed
}

// ControlSystem handles player input.
type ControlSystem struct {
	entities []controlEntity
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
		if vector, changed := e.speed(); changed {
			vector, _ = vector.Normalize()
			vector.MultiplyScalar(dt)
			engo.Mailbox.Dispatch(SpeedMessage{e.BasicEntity, vector})
		}
	}
}

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
