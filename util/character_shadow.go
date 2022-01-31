package util

import (
	"fmt"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

const shadowSpritesheetURL = "spritesheets/LightShadow_pipo.png"

// Shadow entity is an entity shadow.
type Shadow struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceFace
}

func NewCharacterShadow(character *Character) (*Shadow, error) {
	shadowSpritesheet := common.NewSpritesheetFromFile(
		shadowSpritesheetURL, 32, 32,
	)
	if shadowSpritesheet == nil {
		return nil, fmt.Errorf("failed to load shadow spritesheet with url %v", shadowSpritesheetURL)
	}

	shadow := &Shadow{
		BasicEntity: ecs.NewBasic(),
		RenderComponent: common.RenderComponent{
			Drawable:    shadowSpritesheet.Drawable(3),
			StartZIndex: 2,
		},
		SpaceFace: character,
	}

	return shadow, nil
}
