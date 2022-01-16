package character

import (
	"testing"

	"github.com/EngoEngine/ecs"
)

func TestSystem(t *testing.T) {
	t.Parallel()

	character := &Character{}
	system := &CharacterSystem{}

	t.Run("Add", func(t *testing.T) {
		system.Add(
			&character.BasicEntity,
			&character.AnimationComponent,
			&character.SpaceComponent,
			&character.CharacterComponent,
			&character.SpeedComponent,
		)
		want, got := 1, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		system.Remove(character.BasicEntity)
		want, got := 0, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	world := &ecs.World{}
	var characterable *Characterable

	t.Run("AddSystemInterface", func(t *testing.T) {
		world.AddSystemInterface(system, characterable, nil)
		world.AddEntity(character)
		want, got := 1, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	t.Run("Update", func(t *testing.T) {
		system.Update(0)
	})
}
