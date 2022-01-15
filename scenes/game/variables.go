package game

import (
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/systems/action"
	"github.com/eth0net/magicgame/systems/control"
	"github.com/eth0net/magicgame/systems/speed"
	"github.com/eth0net/magicgame/util"
)

// Systems for scene.
var (
	animationSystem = &common.AnimationSystem{}
	collisionSystem = &common.CollisionSystem{Solids: util.CollisionWorld | util.CollisionPlayer | util.CollisionEntity}
	renderSystem    = &common.RenderSystem{}
	actionSystem    = &action.ActionSystem{}
	controlSystem   = &control.ControlSystem{}
	speedSystem     = &speed.SpeedSystem{}
)

// Interfaces for adding entities to Scene systems.
var (
	animationable *common.Animationable
	collisionable *common.Collisionable
	renderable    *common.Renderable
	actionable    *action.Actionable
	controlable   *control.Controlable
	speedable     *speed.Speedable
)
