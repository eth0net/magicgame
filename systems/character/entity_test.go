package character

import (
	"reflect"
	"testing"

	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
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

func TestRunSchedule(t *testing.T) {
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

	char := characterEntity{
		AnimationComponent: &common.AnimationComponent{
			Animations: animations,
		},
		SpaceComponent:     &common.SpaceComponent{},
		CharacterComponent: &CharacterComponent{},
		SpeedComponent:     &speed.SpeedComponent{},
	}

	t.Run("ActStop", func(t *testing.T) {
		char.SpeedComponent.Point = engo.Point{X: 1, Y: 1}
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActStop},
			},
		}
		char.runSchedule(0)

		want, got := engo.Point{}, char.SpeedComponent.Point
		if !want.Equal(got) {
			t.Errorf("wanted Point == %v, got %v", want, got)
		}
	})

	t.Run("ActTurn", func(t *testing.T) {
		char.SpeedComponent.Point = engo.Point{Y: 1}
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActTurn, Point: engo.Point{X: 1}},
			},
		}
		char.AnimationComponent.SelectAnimationByName(AnimationMoveDown)
		char.runSchedule(0)
		char.setAnimation()

		wantPoint, gotPoint := engo.Point{}, char.SpeedComponent.Point
		if !wantPoint.Equal(gotPoint) {
			t.Errorf("wanted Point == %v, got %v", wantPoint, gotPoint)
		}

		wantAnim := animations[AnimationStopRight]
		gotAnim := char.AnimationComponent.CurrentAnimation
		if wantAnim != gotAnim {
			t.Errorf("wanted Animation == %v, got %v", wantAnim, gotAnim)
		}
	})

	t.Run("ActWalk", func(t *testing.T) {
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActWalk, Point: engo.Point{X: 1}},
			},
		}
		t.Skipf("not yet implemented")
	})

	t.Run("ActRun", func(t *testing.T) {
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActRun, Point: engo.Point{X: 1}},
			},
		}
		t.Skipf("not yet implemented")
	})

	t.Run("ActTeleport", func(t *testing.T) {
		char.SpaceComponent.Position = engo.Point{}
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActTeleport, Point: engo.Point{X: 1}},
			},
		}
		char.runSchedule(1)

		want, got := engo.Point{X: 1}, char.SpaceComponent.Position
		if !want.Equal(got) {
			t.Errorf("wanted Point == %v, got %v", want, got)
		}
	})

	t.Run("ActTurnTo", func(t *testing.T) {
		char.SpaceComponent.Position = engo.Point{}
		char.SpeedComponent.Point = engo.Point{Y: 1}
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActTurnTo, Point: engo.Point{X: 1}},
			},
		}
		char.AnimationComponent.SelectAnimationByName(AnimationMoveDown)
		char.runSchedule(0)
		char.setAnimation()

		wantPoint, gotPoint := engo.Point{}, char.SpeedComponent.Point
		if !wantPoint.Equal(gotPoint) {
			t.Errorf("wanted Point == %v, got %v", wantPoint, gotPoint)
		}

		wantAnim := animations[AnimationStopRight]
		gotAnim := char.AnimationComponent.CurrentAnimation
		if wantAnim != gotAnim {
			t.Errorf("wanted Animation == %v, got %v", wantAnim, gotAnim)
		}
	})

	t.Run("ActWalkTo", func(t *testing.T) {
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActWalkTo, Point: engo.Point{X: 1}},
			},
		}
		t.Skipf("not yet implemented")
	})

	t.Run("ActRunTo", func(t *testing.T) {
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActRunTo, Point: engo.Point{X: 1}},
			},
		}
		t.Skipf("not yet implemented")
	})

	t.Run("ActTeleportTo", func(t *testing.T) {
		char.SpaceComponent.Position = engo.Point{}
		char.CharacterComponent.Schedule = Schedule{
			Actions: []Action{
				{Type: ActTeleportTo, Point: engo.Point{X: 1}},
			},
		}
		char.runSchedule(1)

		want, got := engo.Point{X: 1}, char.SpaceComponent.Position
		if !want.Equal(got) {
			t.Errorf("wanted Point == %v, got %v", want, got)
		}
	})
}
