package character

import (
	"reflect"
	"testing"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/raziel2244/magicgame/systems/speed"
)

func TestSetAnimation(t *testing.T) {
	t.Parallel()

	animations := map[string]*common.Animation{
		AnimationMoveUp:    {Name: AnimationMoveUp, Loop: true},
		AnimationMoveDown:  {Name: AnimationMoveDown, Loop: true},
		AnimationMoveLeft:  {Name: AnimationMoveLeft, Loop: true},
		AnimationMoveRight: {Name: AnimationMoveRight, Loop: true},
		AnimationStopUp:    {Name: AnimationStopUp, Loop: true},
		AnimationStopDown:  {Name: AnimationStopDown, Loop: true},
		AnimationStopLeft:  {Name: AnimationStopLeft, Loop: true},
		AnimationStopRight: {Name: AnimationStopRight, Loop: true},
	}

	entity := characterEntity{
		AnimationComponent: &common.AnimationComponent{},
		SpeedComponent:     &speed.SpeedComponent{},
	}
	entity.Animations = animations

	t.Run("SkipNonAnimated", func(t *testing.T) {
		entity.setAnimation()

		var want *common.Animation
		got := entity.CurrentAnimation
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want animation == %v, got %v", want, got)
		}
	})

	t.Run("SkipNonMoving", func(t *testing.T) {
		entity.SelectAnimationByName(AnimationStopDown)
		entity.setAnimation()

		want, got := animations[AnimationStopDown], entity.CurrentAnimation
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want animation == %v, got %v", want, got)
		}
	})

	testCases := []struct {
		Name          string
		AnimationName string
		Speed         engo.Point
	}{
		{"MoveUp", AnimationMoveUp, engo.Point{X: 0, Y: -1}},
		{"StopUp", AnimationStopUp, engo.Point{X: 0, Y: 0}},
		{"MoveDown", AnimationMoveDown, engo.Point{X: 0, Y: 1}},
		{"StopDown", AnimationStopDown, engo.Point{X: 0, Y: 0}},
		{"MoveLeft", AnimationMoveLeft, engo.Point{X: -1, Y: 0}},
		{"StopLeft", AnimationStopLeft, engo.Point{X: 0, Y: 0}},
		{"MoveRight", AnimationMoveRight, engo.Point{X: 1, Y: 0}},
		{"StopRight", AnimationStopRight, engo.Point{X: 0, Y: 0}},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			entity.SpeedComponent.Point = tc.Speed
			entity.setAnimation()

			want, got := animations[tc.AnimationName], entity.CurrentAnimation
			if !reflect.DeepEqual(want, got) {
				t.Errorf("want animation == %v, got %v", want, got)
			}
		})
	}
}
