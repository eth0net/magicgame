package character

import (
	"reflect"
	"testing"

	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/control"
)

func TestNewCharacter(t *testing.T) {
	testCases := []struct {
		Name     string
		Options  NewCharacterOptions
		Expected *Character
	}{
		{
			Name:    "DefaultOptions",
			Options: NewCharacterOptions{},
			Expected: &Character{
				CollisionComponent: common.CollisionComponent{Group: 1},
				ControlComponent: control.ControlComponent{
					Vertical:   control.AxisVertical,
					Horizontal: control.AxisHorizontal,
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {

			character, err := NewCharacter(NewCharacterOptions{})
			if err != nil {
				t.Errorf("expected err == nil, got %v", err)
			}

			t.Run("AnimationComponent", func(t *testing.T) {
				want := tc.Expected.GetAnimationComponent()
				got := character.GetAnimationComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected AnimationComponent == %v, got %v", want, got)
				}
			})

			t.Run("CollisionComponent", func(t *testing.T) {
				want := tc.Expected.GetCollisionComponent()
				got := character.GetCollisionComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected CollisionComponent == %v, got %v", want, got)
				}
			})

			t.Run("RenderComponent", func(t *testing.T) {
				want := tc.Expected.GetRenderComponent()
				got := character.GetRenderComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected RenderComponent == %v, gpt %v", want, got)
				}
			})

			t.Run("SpaceComponent", func(t *testing.T) {
				want := tc.Expected.GetSpaceComponent()
				got := character.GetSpaceComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected SpaceComponent == %v, gpt %v", want, got)
				}
			})

			t.Run("CharacterComponent", func(t *testing.T) {
				want := tc.Expected.GetCharacterComponent()
				got := character.GetCharacterComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected CharacterComponent == %v, gpt %v", want, got)
				}
			})

			t.Run("ControlComponent", func(t *testing.T) {
				want := tc.Expected.GetControlComponent()
				got := character.GetControlComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected ControlComponent == %v, gpt %v", want, got)
				}
			})

			t.Run("SpeedComponent", func(t *testing.T) {
				want := tc.Expected.GetSpeedComponent()
				got := character.GetSpeedComponent()
				if !reflect.DeepEqual(want, got) {
					t.Errorf("expected SpeedComponent == %v, gpt %v", want, got)
				}
			})

		})
	}
}
