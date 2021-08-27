package game

import (
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/character"
	"github.com/eth0net/magicgame/systems/control"
	"github.com/eth0net/magicgame/systems/magic"
	"github.com/eth0net/magicgame/systems/speed"
)

// Systems for scene.
var (
	animationSystem = &common.AnimationSystem{}
	collisionSystem = &common.CollisionSystem{Solids: 1}
	renderSystem    = &common.RenderSystem{}
	characterSystem = &character.CharacterSystem{}
	controlSystem   = &control.ControlSystem{}
	magicSystem     = &magic.System{}
	speedSystem     = &speed.SpeedSystem{}
)

// Interfaces for adding entities to Scene systems.
var (
	animationable *common.Animationable
	collisionable *common.Collisionable
	renderable    *common.Renderable
	characterable *character.Characterable
	controlable   *control.Controlable
	magicable     *magic.Able
	speedable     *speed.Speedable
)
