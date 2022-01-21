package util

import (
	"fmt"
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
	common.RenderComponent
	common.SpaceComponent
}

// An Object within the game world.
type Object struct {
	ecs.BasicEntity
	common.CollisionComponent
	common.SpaceComponent
}

// A Tilemap entity stores the
// data for a single Tiled map.
type Tilemap struct {
	Level *common.Level

	// Tiles contains all tiles from the map.
	Tiles []*Tile

	// Objects contains all objects from the map.
	Objects []*Object

	// Points contains objects of type Point.
	Points map[string]*Object

	// Spawns contains objects of type Spawn.
	Spawns map[string]*Object
}

// NewTilemap constructs a new Tilemap from the provided file url.
func NewTilemap(url string) (tm *Tilemap, err error) {
	resource, err := engo.Files.Resource(url)
	if err != nil {
		return nil, fmt.Errorf("error getting resource: %w", err)
	}

	tm = &Tilemap{
		Level:  resource.(common.TMXResource).Level,
		Points: map[string]*Object{},
		Spawns: map[string]*Object{},
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
				StartZIndex: float32(idx),
			}
			t.SpaceComponent = common.SpaceComponent{
				Position: tile.Point,
				Width:    float32(tm.Level.TileWidth),
				Height:   float32(tm.Level.TileHeight),
				Rotation: tile.Rotation,
			}
			tm.Tiles = append(tm.Tiles, t)
		}
	}

	for _, layer := range tm.Level.ObjectLayers {
		for _, object := range layer.Objects {
			o := &Object{BasicEntity: ecs.NewBasic()}
			o.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{X: object.X, Y: object.Y},
				Width:    object.Width,
				Height:   object.Height,
			}

			shape := common.Shape{}
			for _, tmxLine := range object.Lines {
				for _, line := range tmxLine.Lines {
					l := engo.Line{
						P1: engo.Point{
							X: line.P1.X - object.X,
							Y: line.P1.Y - object.Y,
						},
						P2: engo.Point{
							X: line.P2.X - object.X,
							Y: line.P2.Y - object.Y,
						},
					}
					shape.Lines = append(shape.Lines, l)
				}
			}
			o.SpaceComponent.AddShape(shape)

			switch object.Type {
			case "Collision":
				o.CollisionComponent.Group = CollisionWorld
			case "Point":
				tm.Points[object.Name] = o
			case "Spawn":
				tm.Spawns[object.Name] = o
			}

			tm.Objects = append(tm.Objects, o)
		}
	}

	return tm, nil
}

// AddTilesToWorld adds each Tile entity to the given world.
func (t Tilemap) AddTilesToWorld(w *ecs.World) {
	for _, tile := range t.Tiles {
		w.AddEntity(tile)
	}
	for _, object := range t.Objects {
		w.AddEntity(object)
	}
}
