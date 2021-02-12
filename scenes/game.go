package scenes

import (
	"bytes"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/raziel2244/magicgame/assets"
	"github.com/raziel2244/magicgame/entities"
	"github.com/raziel2244/magicgame/systems"
	"github.com/raziel2244/magicgame/util"
)

const (
	spritesheetURL string = "spritesheets/player.png"
	tilemapURL     string = "tilemaps/fantasy1-min.tmx"
)

var (
	renderSystem = &common.RenderSystem{}
	renderable   *common.Renderable

	animationSystem = &common.AnimationSystem{}
	animationable   *common.Animationable

	characterSystem = &systems.CharacterSystem{}
	characterable   *systems.Characterable

	controlSystem = &systems.ControlSystem{}
	controlable   *systems.Controlable

	speedSystem = &systems.SpeedSystem{}
	speedable   *systems.Speedable
)

// GameSceneType is the unique type identifier for GameScene.
const GameSceneType string = "GameScene"

// GameScene is a playable world scene in the game.
type GameScene struct {
	World   *ecs.World
	Tilemap *util.Tilemap
}

// Preload is called before loading resources.
func (g *GameScene) Preload() {
	files := []string{
		"spritesheets/player.png",
		"tilesets/BaseChip.png",
		"tilesets/Dirt1.png",
		"tilesets/Grass1-Dirt1.png",
		"tilesets/Grass1.png",
		"tilesets/Water1.png",
		"tilemaps/fantasy1-min.tmx",
	}
	for _, file := range files {
		data, err := assets.Asset(file)
		if err != nil {
			log.Fatalf("Unable to locate asset with URL: %v\n", file)
		}
		err = engo.Files.LoadReaderData(file, bytes.NewReader(data))
		if err != nil {
			log.Fatalf("Unable to load asset with URL: %v\n At %v", file, g.Type())
		}
	}
}

// Setup is called before the main loop.
func (g *GameScene) Setup(u engo.Updater) {
	tilemap, err := util.NewTilemap(tilemapURL)
	if err != nil {
		log.Printf(
			"Failed to create Tilemap from resource: %s, error: %s\n",
			tilemapURL, err,
		)
	}
	common.CameraBounds = tilemap.Level.Bounds()
	speedSystem.Level = tilemap.Level

	player, err := entities.NewPlayer(entities.NewPlayerOptions{
		Position:       engo.Point{X: 800, Y: 600},
		SpritesheetURL: spritesheetURL,
		CellWidth:      32,
		CellHeight:     32,
		AnimationRate:  0.1,
		StartZIndex:    3,
	})
	if err != nil {
		log.Printf("Failed to create Player entity, error: %s\n", err)
	}

	entityScroller := &common.EntityScroller{
		SpaceComponent: &player.SpaceComponent,
		TrackingBounds: tilemap.Level.Bounds(),
	}

	g.World, _ = u.(*ecs.World)

	g.World.AddSystemInterface(renderSystem, renderable, nil)
	g.World.AddSystemInterface(animationSystem, animationable, nil)
	g.World.AddSystemInterface(controlSystem, controlable, nil)
	g.World.AddSystemInterface(characterSystem, characterable, nil)
	g.World.AddSystemInterface(speedSystem, speedable, nil)
	g.World.AddSystem(entityScroller)

	tilemap.AddTilesToWorld(g.World)
	g.World.AddEntity(player)
}

// Type returns a unique string representation of GameScene.
func (g *GameScene) Type() string {
	return GameSceneType
}