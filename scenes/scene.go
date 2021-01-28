package scenes

import (
	"github.com/EngoEngine/engo"
)

const (
	// GameSceneType is the unique type identifier for GameScene.
	GameSceneType string = "GameScene"
)

// GameScene is a playable world scene in the game.
type GameScene struct{}

// Preload is called before loading resources
func (g *GameScene) Preload() {}

// Setup is called before the main loop
func (g *GameScene) Setup(u engo.Updater) {}

// Type returns a unique string representation of the Scene, used to identify it
func (g *GameScene) Type() string {
	return GameSceneType
}
