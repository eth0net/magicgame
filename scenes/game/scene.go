package game

import (
	"bytes"
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/assets"
	"github.com/eth0net/magicgame/systems/action"
	"github.com/eth0net/magicgame/util"
)

// Scene is a playable world scene in the game.
type Scene struct {
	World   *ecs.World
	Tilemap *util.Tilemap
}

// Preload is called before loading resources.
func (g *Scene) Preload() {
	files := []string{
		"spritesheets/Male 18-1.png",
		"spritesheets/Female 24-1.png",
		"tilesets/BaseChip.png",
		"tilesets/Dirt1.png",
		"tilesets/Grass1-Dirt1.png",
		"tilesets/Grass1.png",
		"tilesets/Water1.png",
		"tilemaps/fantasy1-min.tmx",
	}
	for _, file := range files {
		data, err := assets.ReadFile(file)
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
func (g *Scene) Setup(u engo.Updater) {
	tilemap, err := util.NewTilemap(tilemapURL)
	if err != nil {
		log.Printf(
			"Failed to create Tilemap from resource: %s, error: %s\n",
			tilemapURL, err,
		)
	}
	speedSystem.Level = tilemap.Level

	playerSpawn, ok := tilemap.Spawns[PlayerSpawnName]
	if !ok {
		panic("no player spawn found in tilemap")
	}

	player, err := util.NewCharacter(util.NewCharacterOptions{
		Position:       playerSpawn.Position,
		SpritesheetURL: spritesheetURL,
		CellWidth:      32,
		CellHeight:     32,
		AnimationRate:  0.1,
		StartZIndex:    3,
	})
	if err != nil {
		log.Printf("Failed to create Player entity, error: %s\n", err)
	}
	player.CollisionComponent.Group = util.CollisionPlayer
	player.ControlComponent.Enabled = true

	npcSpawn, ok := tilemap.Spawns[NPCSpawnName]
	if !ok {
		panic("no npc spawn found in tilemap")
	}

	npc, err := util.NewCharacter(util.NewCharacterOptions{
		Position:       npcSpawn.Position,
		SpritesheetURL: "spritesheets/Female 24-1.png",
		CellWidth:      32,
		CellHeight:     32,
		AnimationRate:  0.1,
		StartZIndex:    3,
	})
	if err != nil {
		log.Printf("Failed to create NPC entity, error: %s\n", err)
	}
	npc.CollisionComponent.Group = util.CollisionEntity
	npc.ActionComponent.Schedule = action.Schedule{
		Actions: []action.Action{
			{Type: action.ActWalkTo, Point: engo.Point{X: 900, Y: 600}},
			{Type: action.ActWalkTo, Point: engo.Point{X: 700, Y: 600}},
		},
		Loop: true,
	}

	entityScroller := &common.EntityScroller{
		SpaceComponent: &player.SpaceComponent,
		TrackingBounds: tilemap.Level.Bounds(),
	}

	g.World, _ = u.(*ecs.World)

	g.World.AddSystemInterface(renderSystem, renderable, nil)
	g.World.AddSystemInterface(animationSystem, animationable, nil)
	g.World.AddSystemInterface(collisionSystem, collisionable, nil)
	g.World.AddSystemInterface(actionSystem, actionable, nil)
	g.World.AddSystemInterface(controlSystem, controlable, nil)
	g.World.AddSystemInterface(speedSystem, speedable, nil)
	g.World.AddSystem(entityScroller)

	tilemap.AddTilesToWorld(g.World)
	g.World.AddEntity(player)
	g.World.AddEntity(npc)

	engo.Mailbox.Listen("WindowResizeMessage", func(msg engo.Message) {
		offsetX, offsetY := engo.GameWidth()/2, engo.GameHeight()/2
		scaleX, scaleY := engo.GetGlobalScale().X, engo.GetGlobalScale().Y

		bounds := tilemap.Level.Bounds()
		bounds.Min.X += offsetX / scaleX
		bounds.Min.Y += offsetY / scaleY
		bounds.Max.X -= offsetX / scaleX
		bounds.Max.Y -= offsetY / scaleY
		common.CameraBounds = bounds
	})
}

// Type returns a unique string representation of Scene.
func (g *Scene) Type() string {
	return SceneType
}
