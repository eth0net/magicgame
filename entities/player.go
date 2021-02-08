package entities

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/raziel2244/magicgame/systems"
)

const (
	spritesheetURL string = "spritesheets/player.png"

	cellWidth  int = 32
	cellHeight int = 32

	animationRate float32 = 0.1
)

// PlayerAnimations contains animations for Player.
var PlayerAnimations = []*common.Animation{
	{Name: systems.AnimationMoveUp, Frames: []int{9, 10, 11}, Loop: true},
	{Name: systems.AnimationMoveDown, Frames: []int{0, 1, 2}, Loop: true},
	{Name: systems.AnimationMoveLeft, Frames: []int{3, 4, 5}, Loop: true},
	{Name: systems.AnimationMoveRight, Frames: []int{6, 7, 8}, Loop: true},
	{Name: systems.AnimationStopUp, Frames: []int{10}, Loop: true},
	{Name: systems.AnimationStopDown, Frames: []int{1}, Loop: true},
	{Name: systems.AnimationStopLeft, Frames: []int{4}, Loop: true},
	{Name: systems.AnimationStopRight, Frames: []int{7}, Loop: true},
}

// Player entity is the player character.
type Player struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.RenderComponent
	common.SpaceComponent
	systems.CharacterComponent
	systems.ControlComponent
	systems.SpeedComponent
}

// NewPlayerOptions provides control
// options when calling NewPlayer.
// type NewPlayerOptions struct {
// 	Position              engo.Point
// 	SpritesheetURL        string
// 	CellWidth, CellHeight int
// 	StartZIndex           float32
// }

// NewPlayer constructs a new Player entity and
// returns it along with any errors encountered.
func NewPlayer() (p *Player, err error) {
	spritesheet := common.NewSpritesheetFromFile(
		spritesheetURL, cellWidth, cellHeight,
	)

	p = &Player{BasicEntity: ecs.NewBasic()}
	p.AnimationComponent = common.NewAnimationComponent(
		spritesheet.Drawables(), animationRate,
	)
	p.SpaceComponent = common.SpaceComponent{
		Width:  float32(cellWidth),
		Height: float32(cellHeight),
	}
	p.RenderComponent = common.RenderComponent{
		Drawable:    spritesheet.Drawable(1),
		StartZIndex: 20,
	}
	p.AnimationComponent.AddAnimations(PlayerAnimations)
	p.AnimationComponent.AddDefaultAnimation(PlayerAnimations[5])

	return p, err
}
