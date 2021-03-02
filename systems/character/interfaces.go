package character

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/speed"
)

// CharacterFace enforces type safe access to underlying CharacterComponent.
type CharacterFace interface {
	GetCharacterComponent() *CharacterComponent
}

// Characterable defines requirements for adding
// entities to the CharacterSystem automatically.
type Characterable interface {
	ecs.BasicFace
	common.AnimationFace
	CharacterFace
	speed.SpeedFace
}
