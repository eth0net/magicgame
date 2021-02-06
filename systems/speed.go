package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// SpeedSystem axes.
const (
	AxisVertical   string  = "vertical"
	AxisHorizontal string  = "horiztonal"
	speedScale     float32 = 6400
)

// SpeedMessageType is the unique type identifier for SpeedMessage.
const SpeedMessageType string = "SpeedMessage"

// SpeedMessage updates the speed of an object within the system.
type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

// Type returns a unique string representation for SpeedMessage.
func (SpeedMessage) Type() string {
	return SpeedMessageType
}

// SpeedComponent stores speed for an entity.
type SpeedComponent struct {
	engo.Point
}

// GetSpeedComponent provides type safe access to SpeedComponent.
func (s *SpeedComponent) GetSpeedComponent() *SpeedComponent {
	return s
}

type speedEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*SpeedComponent
}

// SpeedSystem handles speed and position updates for entities.
type SpeedSystem struct {
	entities []speedEntity
	Level    *common.Level
}

// New initialises SpeedSystem when it's added to the world.
func (ss *SpeedSystem) New(*ecs.World) {
	engo.Input.RegisterButton(upButton, engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton(downButton, engo.KeyS, engo.KeyArrowDown)
	engo.Input.RegisterButton(leftButton, engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton(rightButton, engo.KeyD, engo.KeyArrowRight)

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

	engo.Mailbox.Listen(SpeedMessageType, func(msg engo.Message) {
		speed, ok := msg.(SpeedMessage)
		if !ok {
			return
		}

		for _, e := range ss.entities {
			if e.ID() != speed.ID() {
				continue
			}
			e.Point = speed.Point
		}
	})
}

// Add an entity to the SpeedSystem.
func (ss *SpeedSystem) Add(
	basic *ecs.BasicEntity,
	space *common.SpaceComponent,
	speed *SpeedComponent,
) {
	if ss.entities == nil {
		ss.entities = []speedEntity{}
	}
	ss.entities = append(ss.entities, speedEntity{basic, space, speed})
}

// AddByInterface adds entities to the system via Speedable interface.
func (ss *SpeedSystem) AddByInterface(i ecs.Identifier) {
	e := i.(Speedable)
	ss.Add(
		e.GetBasicEntity(),
		e.GetSpaceComponent(),
		e.GetSpeedComponent(),
	)
}

// Remove an entity from the SpeedSystem.
func (ss *SpeedSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range ss.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		ss.entities = append(ss.entities[:del], ss.entities[del+1:]...)
	}
}

// Update the SpeedSystem this frame.
func (ss *SpeedSystem) Update(dt float32) {
	speedX := speedScale * dt * engo.GetGlobalScale().X
	speedY := speedScale * dt * engo.GetGlobalScale().Y
	for _, e := range ss.entities {
		e.Position.X = e.Position.X + speedX*e.SpeedComponent.X
		e.Position.Y = e.Position.Y + speedY*e.SpeedComponent.Y

		// limit to map borders
		var limitX float32 = ss.Level.Bounds().Max.X - e.Width
		var limitY float32 = ss.Level.Bounds().Max.Y - e.Height
		switch {
		case e.SpaceComponent.Position.X < 0:
			e.SpaceComponent.Position.X = 0
		case e.SpaceComponent.Position.X > limitX:
			e.SpaceComponent.Position.X = limitX
		}
		switch {
		case e.SpaceComponent.Position.Y < 0:
			e.SpaceComponent.Position.Y = 0
		case e.SpaceComponent.Position.Y > limitY:
			e.SpaceComponent.Position.Y = limitY
		}
	}
}

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
