package character

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

type characterEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*CharacterComponent
	*speed.SpeedComponent
}

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
