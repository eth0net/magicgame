package control

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
)

type controlEntity struct {
	*ecs.BasicEntity
	*ControlComponent
}

func (ce *controlEntity) speed() (p engo.Point, changed bool) {
	var (
		upReleased    = engo.Input.Button(ButtonUp).JustReleased()
		downReleased  = engo.Input.Button(ButtonDown).JustReleased()
		leftReleased  = engo.Input.Button(ButtonLeft).JustReleased()
		rightReleased = engo.Input.Button(ButtonRight).JustReleased()
		upHeld        = engo.Input.Button(ButtonUp).Down()
		downHeld      = engo.Input.Button(ButtonDown).Down()
		leftHeld      = engo.Input.Button(ButtonLeft).Down()
		rightHeld     = engo.Input.Button(ButtonRight).Down()
	)

	oldX := engo.Input.Axis(ce.Horizontal).Value()
	oldY := engo.Input.Axis(ce.Vertical).Value()

	p.X, p.Y = oldX, oldY

	switch {
	case upReleased, downReleased:
		changed = true // TODO: try removing
		p.Y = 0
	case leftReleased, rightReleased:
		changed = true // TODO: try removing
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
