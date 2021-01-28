package scenes

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// GameSceneType is the unique type identifier for GameScene.
	GameSceneType string = "GameScene"

	// camera config
	cameraScrollSpeed float32 = 400
	cameraZoomSpeed   float32 = -0.125
)

// A Tile entity stores the contents
// of a single tile of the game map.
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// GameScene is a playable world scene in the game.
type GameScene struct{}

// Preload is called before loading resources
func (g *GameScene) Preload() {
	err := engo.Files.Load("tilemap/map-ex.tmx")
	if err != nil {
		panic(err)
	}
}

// Setup is called before the main loop
func (g *GameScene) Setup(u engo.Updater) {
	resource, err := engo.Files.Resource("tilemap/map-ex.tmx")
	if err != nil {
		panic(err)
	}
	level := resource.(common.TMXResource).Level
	common.CameraBounds = level.Bounds()

	tiles := []*Tile{}
	for idx, tileLayer := range level.TileLayers {
		for _, tileElement := range tileLayer.Tiles {
			if tileElement.Image == nil {
				log.Printf("Tile is lacking image at point: %v", tileElement.Point)
			}
			tile := &Tile{BasicEntity: ecs.NewBasic()}
			tile.RenderComponent = common.RenderComponent{
				Drawable:    tileElement.Image,
				Scale:       engo.Point{X: 1, Y: 1},
				StartZIndex: float32(idx),
			}
			tile.Position = tileElement.Point
			tiles = append(tiles, tile)
		}
	}

	keyboardScroller := common.NewKeyboardScroller(
		cameraScrollSpeed,
		engo.DefaultHorizontalAxis,
		engo.DefaultVerticalAxis,
	)

	mouseZoomer := &common.MouseZoomer{
		ZoomSpeed: cameraZoomSpeed,
	}

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(keyboardScroller)
	world.AddSystem(mouseZoomer)

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.RenderSystem:
			for _, tile := range tiles {
				sys.Add(&tile.BasicEntity, &tile.RenderComponent, &tile.SpaceComponent)
			}
		}
	}
}

// Type returns a unique string representation of the Scene, used to identify it
func (g *GameScene) Type() string {
	return GameSceneType
}
