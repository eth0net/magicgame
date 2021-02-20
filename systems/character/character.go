package character

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/raziel2244/magicgame/systems/control"
	"github.com/raziel2244/magicgame/systems/speed"
)

// CharacterAnimations contains animations for Character.
var CharacterAnimations = []*common.Animation{
	{Name: AnimationMoveUp, Frames: []int{9, 10, 11}, Loop: true},
	{Name: AnimationMoveDown, Frames: []int{0, 1, 2}, Loop: true},
	{Name: AnimationMoveLeft, Frames: []int{3, 4, 5}, Loop: true},
	{Name: AnimationMoveRight, Frames: []int{6, 7, 8}, Loop: true},
	{Name: AnimationStopUp, Frames: []int{10}, Loop: true},
	{Name: AnimationStopDown, Frames: []int{1}, Loop: true},
	{Name: AnimationStopLeft, Frames: []int{4}, Loop: true},
	{Name: AnimationStopRight, Frames: []int{7}, Loop: true},
}

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
	spritesheet := common.NewSpritesheetFromFile(
		o.SpritesheetURL, o.CellWidth, o.CellHeight,
	)
	if spritesheet == nil {
		err = fmt.Errorf("Failed to load spritesheet with url %v", o.SpritesheetURL)
		return p, err
	}

	p = &Character{BasicEntity: ecs.NewBasic()}
	p.AnimationComponent = common.NewAnimationComponent(
		spritesheet.Drawables(), o.AnimationRate,
	)
	p.CollisionComponent.Group = 1
	p.SpaceComponent = common.SpaceComponent{
		Position: o.Position,
		Width:    float32(o.CellWidth),
		Height:   float32(o.CellHeight),
	}
	p.RenderComponent = common.RenderComponent{
		Drawable:    spritesheet.Drawable(1),
		StartZIndex: o.StartZIndex,
	}
	p.ControlComponent = control.ControlComponent{
		Vertical:   control.AxisVertical,
		Horizontal: control.AxisHorizontal,
	}
	p.AnimationComponent.AddAnimations(CharacterAnimations)
	p.AnimationComponent.AddDefaultAnimation(CharacterAnimations[5])

	return p, err
}
