package fireball

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

const SpritesheetURL string = "effects/fireball.png"

type Entity struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
	speed.SpeedComponent
}

func New(position, speed engo.Point) *Entity {
	e := &Entity{BasicEntity: ecs.NewBasic()}

	spritesheet := common.NewSpritesheetFromFile(SpritesheetURL, 64, 64)

	frames := make([]int, spritesheet.CellCount())
	for i := 0; i < spritesheet.CellCount(); i++ {
		frames[i] = i
	}

	e.AnimationComponent = common.NewAnimationComponent(spritesheet.Drawables(), .1)
	e.AnimationComponent.AddDefaultAnimation(&common.Animation{Name: "default", Frames: frames, Loop: true})

	e.RenderComponent = common.RenderComponent{
		Drawable:    spritesheet.Drawable(0),
		StartZIndex: 5,
	}

	e.CollisionComponent.Group = 1

	e.SpaceComponent = common.SpaceComponent{
		Position: position,
		Width:    32,
		Height:   32,
	}

	e.SpeedComponent.Point = speed

	return e
}
