package util

import (
	"fmt"
	"log"
	"strconv"

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
	Level   *common.Level
	Tiles   []*Tile
	Objects []*Object
}

// NewTilemap constructs a new Tilemap from the provided file url.
func NewTilemap(url string) (tm *Tilemap, err error) {
	resource, err := engo.Files.Resource(url)
	if err != nil {
		return nil, fmt.Errorf("error getting resource: %w", err)
	}

	tm = &Tilemap{Level: resource.(common.TMXResource).Level}

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

	for _, layer := range tm.Level.ObjectLayers {
		for _, object := range layer.Objects {
			o := &Object{BasicEntity: ecs.NewBasic()}
			o.SpaceComponent = common.SpaceComponent{
				Position: engo.Point{X: object.X, Y: object.Y},
				Width:    32,
				Height:   32,
			}

			var collision bool
			for _, property := range layer.Properties {
				if property.Name != "Collision" {
					continue
				}
				collision, _ = strconv.ParseBool(property.Value)
			}
			if collision {
				s := common.Shape{}
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
						s.Lines = append(s.Lines, l)
					}
				}
				o.AddShape(s)
				o.CollisionComponent.Group = CollisionWorld
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
