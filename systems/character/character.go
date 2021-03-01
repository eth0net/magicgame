package character

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/control"
	"github.com/eth0net/magicgame/systems/speed"
)

// Character entity is a game character.
type Character struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
	CharacterComponent
	control.ControlComponent
	speed.SpeedComponent
}

// NewCharacterOptions provides control
// options when calling NewCharacter.
type NewCharacterOptions struct {
	Position              engo.Point
	SpritesheetURL        string
	CellWidth, CellHeight int
	AnimationRate         float32
	StartZIndex           float32
}

// NewCharacter constructs a new Character entity and
// returns it along with any errors encountered.
func NewCharacter(o NewCharacterOptions) (p *Character, err error) {
	p = &Character{BasicEntity: ecs.NewBasic()}

	if o.SpritesheetURL != "" {
		spritesheet := common.NewSpritesheetFromFile(
			o.SpritesheetURL, o.CellWidth, o.CellHeight,
		)
		if spritesheet == nil {
			err = fmt.Errorf("Failed to load spritesheet with url %v", o.SpritesheetURL)
			return p, err
		}
		p.AnimationComponent = common.NewAnimationComponent(
			spritesheet.Drawables(), o.AnimationRate,
		)
		p.RenderComponent = common.RenderComponent{
			Drawable:    spritesheet.Drawable(1),
			StartZIndex: o.StartZIndex,
		}
		p.AnimationComponent.AddAnimations(CharacterAnimations)
		p.AnimationComponent.AddDefaultAnimation(CharacterAnimations[5])
	}

	p.CollisionComponent.Group = 1
	p.SpaceComponent = common.SpaceComponent{
		Position: o.Position,
		Width:    float32(o.CellWidth),
		Height:   float32(o.CellHeight),
	}
	p.ControlComponent = control.ControlComponent{
		Vertical:   control.AxisVertical,
		Horizontal: control.AxisHorizontal,
	}

	return p, err
}
