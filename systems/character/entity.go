package character

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

type characterEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*common.SpaceComponent
	*CharacterComponent
	*speed.SpeedComponent
}

// setAnimation sets the characterEntity Animation
// using the current state of the SpeedComponent.
func (ce *characterEntity) setAnimation() {
	point := ce.GetSpeedComponent().Point
	currentAnimation := ce.AnimationComponent.CurrentAnimation
	if currentAnimation == nil {
		return
	}
	newAnimationName := currentAnimation.Name

	var (
		xIsNegative bool = point.X < 0
		xIsPositive bool = point.X > 0
		xIsZero     bool = point.X == 0

		yIsNegative bool = point.Y < 0
		yIsPositive bool = point.Y > 0
		yIsZero     bool = point.Y == 0
	)

	switch {
	case xIsZero && yIsZero:
		switch currentAnimation.Name {
		case AnimationMoveUp:
			newAnimationName = AnimationStopUp
		case AnimationMoveDown:
			newAnimationName = AnimationStopDown
		case AnimationMoveLeft:
			newAnimationName = AnimationStopLeft
		case AnimationMoveRight:
			newAnimationName = AnimationStopRight
		}
	case xIsZero && yIsNegative:
		newAnimationName = AnimationMoveUp
	case xIsZero && yIsPositive:
		newAnimationName = AnimationMoveDown
	case xIsNegative:
		newAnimationName = AnimationMoveLeft
	case xIsPositive:
		newAnimationName = AnimationMoveRight
	}

	if currentAnimation.Name != newAnimationName {
		ce.SelectAnimationByName(newAnimationName)
	}
}

// runSchedule executes code on the characterEntity
// according to the current Action for the Schedule
// and then updates the Schedule accordingly.
func (ce *characterEntity) runSchedule(dt float32) {
	actionPtr := ce.CurrentAction()
	if actionPtr == nil {
		return
	}

	action := *actionPtr

	var actComplete bool

	switch action.Type {
	case ActStop:
		ce.SpeedComponent.Point = engo.Point{}
		actComplete = true
		break

	case ActTurn:
		ce.SpeedComponent.Point = action.Point
		ce.setAnimation()
		ce.SpeedComponent.Point = engo.Point{}
		ce.setAnimation()
		actComplete = true
		break

	case ActWalk:
		vector, _ := action.Point.Normalize()
		vector.MultiplyScalar(dt)

		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		actComplete = true
		break

	case ActRun:
		vector, _ := action.Point.Normalize()
		vector.MultiplyScalar(dt * 2)

		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		actComplete = true
		break

	case ActTeleport:
		vector := action.Point.MultiplyScalar(dt)
		ce.SpaceComponent.Position.Add(*vector)
		actComplete = true
		break

	case ActTurnTo:
		direction := action.Point.Subtract(ce.SpaceComponent.Position)
		ce.SpeedComponent.Point = *direction
		ce.setAnimation()
		ce.SpeedComponent.Point = engo.Point{}
		ce.setAnimation()
		actComplete = ce.SpeedComponent.Point == action.Point
		break

	case ActWalkTo:
		vector := action.Point
		vector.Subtract(ce.SpaceComponent.Position)
		vector, _ = vector.Normalize()
		vector.MultiplyScalar(dt)
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		if ce.SpaceComponent.Position.Equal(action.Point) {
			actComplete = true
		}
		break

	case ActRunTo:
		direction := action.Point.Subtract(ce.SpaceComponent.Position)
		vector, _ := direction.Normalize()
		vector.MultiplyScalar(dt * 2)
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		if ce.SpaceComponent.Position.Equal(action.Point) {
			actComplete = true
		}
		break

	case ActTeleportTo:
		ce.SpaceComponent.Position = action.Point
		actComplete = true
		break
	}

	ce.currentDuration += dt
	durationPassed := ce.currentDuration >= action.Duration
	if actComplete && durationPassed {
		ce.StepAction()
	}
}
