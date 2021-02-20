package game

import (
	"github.com/EngoEngine/engo/common"
	"github.com/raziel2244/magicgame/systems/character"
	"github.com/raziel2244/magicgame/systems/control"
	"github.com/raziel2244/magicgame/systems/speed"
)

// Systems for scene.
var (
	animationSystem = &common.AnimationSystem{}
	collisionSystem = &common.CollisionSystem{Solids: 1}
	renderSystem    = &common.RenderSystem{}
	characterSystem = &character.CharacterSystem{}
	controlSystem   = &control.ControlSystem{}
	speedSystem     = &speed.SpeedSystem{}
)

// Interfaces for adding entities to Scene systems.
var (
	animationable *common.Animationable
	collisionable *common.Collisionable
	renderable    *common.Renderable
	characterable *character.Characterable
	controlable   *control.Controlable
	speedable     *speed.Speedable
)
