package util

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

// A Tile entity stores the contents
// of a single tile of the game map.
type Tile struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.CollisionComponent
	common.RenderComponent
	common.SpaceComponent
}

// A Tilemap entity stores the
// data for a single Tiled map.
type Tilemap struct {
	Level *common.Level
	Tiles []*Tile
}

// NewTilemap constructs a new Tilemap from the provided file url.
func NewTilemap(url string) (tm *Tilemap, err error) {
	resource, err := engo.Files.Resource(url)
	if err != nil {
		return nil, err
	}

	tm = &Tilemap{
		Level: resource.(common.TMXResource).Level,
		Tiles: []*Tile{},
	}

	for idx, layer := range tm.Level.TileLayers {
		for _, tile := range layer.Tiles {
			if tile.Image == nil {
				log.Printf("Tile is lacking image at point: %v", tile.Point)
			}
			t := &Tile{BasicEntity: ecs.NewBasic()}
			if len(tile.Drawables) > 0 {
				t.AnimationComponent = common.NewAnimationComponent(tile.Drawables, 0.1)
				t.AnimationComponent.AddDefaultAnimation(tile.Animation)
			}
			t.RenderComponent = common.RenderComponent{
				Drawable:    tile.Image,
				Scale:       engo.Point{X: 1, Y: 1},
				StartZIndex: float32(idx),
			}
			t.SpaceComponent = common.SpaceComponent{
				Position: tile.Point,
			}
			tm.Tiles = append(tm.Tiles, t)
		}
	}

	return tm, err
}

// AddTilesToWorld adds each Tile entity to the given world.
func (t Tilemap) AddTilesToWorld(w *ecs.World) {
	for _, tile := range t.Tiles {
		w.AddEntity(tile)
	}
}
