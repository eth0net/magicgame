package action

import (
	"testing"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

func TestSystem(t *testing.T) {
	t.Parallel()

	entity := &struct {
		ecs.BasicEntity
		common.AnimationComponent
		common.SpaceComponent
		ActionComponent
		speed.SpeedComponent
	}{
		BasicEntity: ecs.NewBasic(),
	}
	system := &ActionSystem{}

	t.Run("Add", func(t *testing.T) {
		system.Add(
			entity.GetBasicEntity(),
			entity.GetAnimationComponent(),
			entity.GetSpaceComponent(),
			entity.GetActionComponent(),
			entity.GetSpeedComponent(),
		)
		want, got := 1, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	t.Run("Remove", func(t *testing.T) {
		system.Remove(entity.BasicEntity)
		want, got := 0, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	t.Run("AddSystemInterface", func(t *testing.T) {
		var actionable *Actionable
		var world ecs.World
		world.AddSystemInterface(system, actionable, nil)
		world.AddEntity(entity)
		want, got := 1, len(system.entities)
		if want != got {
			t.Errorf("want len == %v, got %v", want, got)
		}
	})

	t.Run("Update", func(t *testing.T) {
		system.Update(0)
	})
}
