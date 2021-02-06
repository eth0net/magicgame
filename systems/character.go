package systems

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo/common"
)

// Animation name constants.
const (
	AnimationMoveUp    string = "moveUp"
	AnimationMoveDown  string = "moveDown"
	AnimationMoveLeft  string = "moveLeft"
	AnimationMoveRight string = "moveRight"
	AnimationStopUp    string = "stopUp"
	AnimationStopDown  string = "stopDown"
	AnimationStopLeft  string = "stopLeft"
	AnimationStopRight string = "stopRight"
)

// CharacterComponent marks entities for use with a CharacterSystem.
type CharacterComponent struct{}

// GetCharacterComponent provides type safe access to CharacterComponent.
func (pc *CharacterComponent) GetCharacterComponent() *CharacterComponent {
	return pc
}

type characterEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*CharacterComponent
	*SpeedComponent
}

func (ce *characterEntity) setAnimation() {
	point := ce.GetSpeedComponent().Point
	currentAnimation := ce.AnimationComponent.CurrentAnimation
	if currentAnimation == nil {
		return
	}
	newAnimationName := currentAnimation.Name

	var (
		xIsNegative bool = point.X < 0
		xIsPositive bool = point.X > 0
		xIsZero     bool = point.X == 0

		yIsNegative bool = point.Y < 0
		yIsPositive bool = point.Y > 0
		yIsZero     bool = point.Y == 0
	)

	switch {
	case xIsZero && yIsZero:
		switch currentAnimation.Name {
		case AnimationMoveUp:
			newAnimationName = AnimationStopUp
		case AnimationMoveDown:
			newAnimationName = AnimationStopDown
		case AnimationMoveLeft:
			newAnimationName = AnimationStopLeft
		case AnimationMoveRight:
			newAnimationName = AnimationStopRight
		}
	case xIsZero && yIsNegative:
		newAnimationName = AnimationMoveUp
	case xIsZero && yIsPositive:
		newAnimationName = AnimationMoveDown
	case xIsNegative:
		newAnimationName = AnimationMoveLeft
	case xIsPositive:
		newAnimationName = AnimationMoveRight
	}

	if currentAnimation.Name != newAnimationName {
		ce.SelectAnimationByName(newAnimationName)
	}

}

// CharacterSystem handles character-specific operations.
type CharacterSystem struct {
	entities []characterEntity
}

// Add an entity to the CharacterSystem.
func (cs *CharacterSystem) Add(
	b *ecs.BasicEntity,
	a *common.AnimationComponent,
	c *CharacterComponent,
	s *SpeedComponent,
) {
	if cs.entities == nil {
		cs.entities = []characterEntity{}
	}
	cs.entities = append(cs.entities, characterEntity{b, a, c, s})
}

// AddByInterface adds entities to the system via Characterable interface.
func (cs *CharacterSystem) AddByInterface(i ecs.Identifier) {
	e := i.(Characterable)
	cs.Add(
		e.GetBasicEntity(),
		e.GetAnimationComponent(),
		e.GetCharacterComponent(),
		e.GetSpeedComponent(),
	)
}

// Remove an entity from the CharacterSystem.
func (cs *CharacterSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range cs.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		cs.entities = append(cs.entities[:del], cs.entities[del+1:]...)
	}
}

// Update the CharacterSystem this frame.
func (cs *CharacterSystem) Update(dt float32) {
	for _, e := range cs.entities {
		e.setAnimation()
	}
}

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
	SpeedFace
}
