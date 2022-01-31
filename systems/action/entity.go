package action

import (
	"math"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

type actionEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*common.SpaceComponent
	*ActionComponent
	*speed.SpeedComponent
}

// setAnimation sets the actionEntity Animation
// using the current state of the SpeedComponent.
func (ce *actionEntity) setAnimation() {
	point := ce.GetSpeedComponent().Point
	currentAnimation := ce.AnimationComponent.CurrentAnimation
	if currentAnimation == nil {
		return
	}
	newAnimationName := currentAnimation.Name

	var (
		xIsNegative bool = math.Round((float64(point.X)*100)/100) < 0
		xIsPositive bool = math.Round((float64(point.X)*100)/100) > 0
		xIsZero     bool = !xIsNegative && !xIsPositive

		yIsNegative bool = math.Round((float64(point.Y)*100)/100) < 0
		yIsPositive bool = math.Round((float64(point.Y)*100)/100) > 0
		yIsZero     bool = !yIsNegative && !yIsPositive

		yIsBigger bool = math.Abs(float64(point.Y)) > math.Abs(float64(point.X))
	)

	switch {
	case xIsNegative:
		newAnimationName = AnimationMoveLeft
	case xIsPositive:
		newAnimationName = AnimationMoveRight
	case yIsBigger && yIsNegative:
		newAnimationName = AnimationMoveUp
	case yIsBigger && yIsPositive:
		newAnimationName = AnimationMoveDown
	case xIsZero && yIsZero:
		switch currentAnimation.Name {
		case AnimationMoveLeft:
			newAnimationName = AnimationStopLeft
		case AnimationMoveRight:
			newAnimationName = AnimationStopRight
		case AnimationMoveUp:
			newAnimationName = AnimationStopUp
		case AnimationMoveDown:
			newAnimationName = AnimationStopDown
		}
	}

	if currentAnimation.Name != newAnimationName {
		ce.SelectAnimationByName(newAnimationName)
	}
}

// runSchedule executes code on the actionEntity
// according to the current Action for the Schedule
// and then updates the Schedule accordingly.
func (ce *actionEntity) runSchedule(dt float32) {
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

	case ActTurn:
		ce.SpeedComponent.Point = action.Point
		ce.setAnimation()
		ce.SpeedComponent.Point = engo.Point{}
		ce.setAnimation()
		actComplete = true

	case ActWalk:
		vector, _ := action.Point.Normalize()
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		actComplete = true

	case ActRun:
		vector, _ := action.Point.Normalize()
		vector.MultiplyScalar(2)
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		actComplete = true

	case ActTeleport:
		vector := action.Point.MultiplyScalar(dt)
		ce.SpaceComponent.Position.Add(*vector)
		actComplete = true

	case ActTurnTo:
		direction := action.Point.Subtract(ce.SpaceComponent.Position)
		ce.SpeedComponent.Point = *direction
		ce.setAnimation()
		ce.SpeedComponent.Point = engo.Point{}
		ce.setAnimation()
		actComplete = ce.SpeedComponent.Point == action.Point

	case ActWalkTo:
		vector := action.Point
		vector.Subtract(ce.SpaceComponent.Position)
		vector, _ = vector.Normalize()
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		if ce.SpaceComponent.Position.Equal(action.Point) {
			actComplete = true
		}

	case ActRunTo:
		direction := action.Point.Subtract(ce.SpaceComponent.Position)
		vector, _ := direction.Normalize()
		vector.MultiplyScalar(2)
		engo.Mailbox.Dispatch(speed.SpeedMessage{
			BasicEntity: ce.BasicEntity,
			Point:       vector,
		})
		if ce.SpaceComponent.Position.Equal(action.Point) {
			actComplete = true
		}

	case ActTeleportTo:
		ce.SpaceComponent.Position = action.Point
		actComplete = true
	}

	ce.currentDuration += dt
	durationPassed := ce.currentDuration >= action.Duration
	if actComplete && durationPassed {
		ce.StepAction()
	}
}
