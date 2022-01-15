package speed

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// SpeedSystem handles speed and position updates for entities.
type SpeedSystem struct {
	entities []speedEntity
	Level    *common.Level
}

// New initialises SpeedSystem when it's added to the world.
func (ss *SpeedSystem) New(*ecs.World) {
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
	v := engo.GetGlobalScale()
	v.MultiplyScalar(dt * speedScale)

	for _, e := range ss.entities {
		e.Position.X = e.Position.X + v.X*e.SpeedComponent.X
		e.Position.Y = e.Position.Y + v.Y*e.SpeedComponent.Y

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
