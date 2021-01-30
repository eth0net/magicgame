package scenes

import (
	"log"

	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
)

const (
	// camera config
	cameraScrollSpeed float32 = 400
	cameraZoomSpeed   float32 = -0.125

	cellWidth  int = 32
	cellHeight int = 32

	heroSpritesheetURL string = "spritesheets/hero.png"
	tilemapURL         string = "tilemaps/fantasy1-min.tmx"

	axisVertical   string = "vertical"
	axisHorizontal string = "horizontal"

	speedScale float32 = 32
)

// Animations for character sprites.
var (
	WalkUpAnimation    *common.Animation
	WalkDownAnimation  *common.Animation
	WalkLeftAnimation  *common.Animation
	WalkRightAnimation *common.Animation
	StopUpAnimation    *common.Animation
	StopDownAnimation  *common.Animation
	StopLeftAnimation  *common.Animation
	StopRightAnimation *common.Animation
	animations         []*common.Animation
)

var (
	upButton    string = "up"
	downButton  string = "down"
	leftButton  string = "left"
	rightButton string = "right"
)

var (
	levelWidth  float32
	levelHeight float32
)

// Hero entity is the player character.
type Hero struct {
	ecs.BasicEntity
	common.AnimationComponent
	common.RenderComponent
	common.SpaceComponent
	ControlComponent
	SpeedComponent
}

// A Tile entity stores the contents
// of a single tile of the game map.
type Tile struct {
	ecs.BasicEntity
	common.RenderComponent
	common.SpaceComponent
}

// GameSceneType is the unique type identifier for GameScene.
const GameSceneType string = "GameScene"

// GameScene is a playable world scene in the game.
type GameScene struct{}

// Preload is called before loading resources
func (g *GameScene) Preload() {
	err := engo.Files.Load(tilemapURL, heroSpritesheetURL)
	if err != nil {
		panic(err)
	}

	WalkUpAnimation = &common.Animation{
		Name:   "walkUp",
		Frames: []int{9, 10, 11},
		Loop:   true,
	}

	WalkDownAnimation = &common.Animation{
		Name:   "walkDown",
		Frames: []int{0, 1, 2},
		Loop:   true,
	}

	WalkLeftAnimation = &common.Animation{
		Name:   "walkLeft",
		Frames: []int{3, 4, 5},
		Loop:   true,
	}

	WalkRightAnimation = &common.Animation{
		Name:   "walkRight",
		Frames: []int{6, 7, 8},
		Loop:   true,
	}

	StopUpAnimation = &common.Animation{
		Name:   "stopUp",
		Frames: []int{10},
	}

	StopDownAnimation = &common.Animation{
		Name:   "stopDown",
		Frames: []int{1},
	}

	StopLeftAnimation = &common.Animation{
		Name:   "stopLeft",
		Frames: []int{4},
	}

	StopRightAnimation = &common.Animation{
		Name:   "stopRight",
		Frames: []int{7},
	}

	animations = []*common.Animation{
		WalkUpAnimation,
		WalkDownAnimation,
		WalkLeftAnimation,
		WalkRightAnimation,
		StopUpAnimation,
		StopDownAnimation,
		StopLeftAnimation,
		StopRightAnimation,
	}

	engo.Input.RegisterButton(upButton, engo.KeyW, engo.KeyArrowUp)
	engo.Input.RegisterButton(downButton, engo.KeyS, engo.KeyArrowDown)
	engo.Input.RegisterButton(leftButton, engo.KeyA, engo.KeyArrowLeft)
	engo.Input.RegisterButton(rightButton, engo.KeyD, engo.KeyArrowRight)
}

// Setup is called before the main loop
func (g *GameScene) Setup(u engo.Updater) {
	resource, err := engo.Files.Resource(tilemapURL)
	if err != nil {
		panic(err)
	}
	level := resource.(common.TMXResource).Level
	common.CameraBounds = level.Bounds()

	levelWidth = level.Bounds().Max.X
	levelHeight = level.Bounds().Max.Y

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

	spritesheet := common.NewSpritesheetFromFile(
		heroSpritesheetURL,
		cellWidth,
		cellHeight,
	)
	hero := &Hero{BasicEntity: ecs.NewBasic()}
	hero.AnimationComponent = common.NewAnimationComponent(
		spritesheet.Drawables(), 0.1,
	)
	hero.Drawable = spritesheet.Cell(1)
	hero.SpaceComponent = common.SpaceComponent{
		Position: engo.Point{
			X: levelWidth / 2,
			Y: levelHeight/2 - 192,
		},
		Width:  float32(cellWidth),
		Height: float32(cellHeight),
	}
	hero.ControlComponent = ControlComponent{
		Vertical:   axisVertical,
		Horizontal: axisHorizontal,
	}
	hero.SpeedComponent = SpeedComponent{}
	hero.AddAnimations(animations)
	hero.SelectAnimationByAction(StopDownAnimation)
	hero.SetZIndex(2)

	entityScroller := &common.EntityScroller{
		SpaceComponent: &hero.SpaceComponent,
		TrackingBounds: level.Bounds(),
	}

	world, _ := u.(*ecs.World)
	world.AddSystem(&common.RenderSystem{})
	world.AddSystem(&common.AnimationSystem{})
	world.AddSystem(&ControlSystem{})
	world.AddSystem(&SpeedSystem{})
	world.AddSystem(entityScroller)

	for _, system := range world.Systems() {
		switch sys := system.(type) {
		case *common.AnimationSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.AnimationComponent,
				&hero.RenderComponent,
			)

		case *common.RenderSystem:
			sys.Add(&hero.BasicEntity, &hero.RenderComponent, &hero.SpaceComponent)
			for _, tile := range tiles {
				sys.Add(&tile.BasicEntity, &tile.RenderComponent, &tile.SpaceComponent)
			}

		case *ControlSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.AnimationComponent,
				&hero.SpaceComponent,
				&hero.ControlComponent,
			)

		case *SpeedSystem:
			sys.Add(
				&hero.BasicEntity,
				&hero.SpaceComponent,
				&hero.SpeedComponent,
			)
		}
	}

	engo.Input.RegisterAxis(
		axisVertical,
		engo.AxisKeyPair{Min: engo.KeyW, Max: engo.KeyS},
		engo.AxisKeyPair{Min: engo.KeyArrowUp, Max: engo.KeyArrowDown},
	)

	engo.Input.RegisterAxis(
		axisHorizontal,
		engo.AxisKeyPair{Min: engo.KeyA, Max: engo.KeyD},
		engo.AxisKeyPair{Min: engo.KeyArrowLeft, Max: engo.KeyArrowRight},
	)
}

// Type returns a unique string representation of the Scene, used to identify it
func (g *GameScene) Type() string {
	return GameSceneType
}

// ControlComponent stores control input for player entity.
type ControlComponent struct {
	Vertical, Horizontal string
}

type controlEntity struct {
	*ecs.BasicEntity
	*common.AnimationComponent
	*common.SpaceComponent
	*ControlComponent
}

func (ce *controlEntity) speed() (p engo.Point, changed bool) {
	var (
		upReleased    = engo.Input.Button(upButton).JustReleased()
		downReleased  = engo.Input.Button(downButton).JustReleased()
		leftReleased  = engo.Input.Button(leftButton).JustReleased()
		rightReleased = engo.Input.Button(rightButton).JustReleased()
		upHeld        = engo.Input.Button(upButton).Down()
		downHeld      = engo.Input.Button(downButton).Down()
		leftHeld      = engo.Input.Button(leftButton).Down()
		rightHeld     = engo.Input.Button(rightButton).Down()
	)

	oldX := engo.Input.Axis(ce.Horizontal).Value()
	oldY := engo.Input.Axis(ce.Vertical).Value()

	p.X, p.Y = oldX, oldY

	switch {
	case upReleased, downReleased:
		changed = true
		p.Y = 0
	case leftReleased, rightReleased:
		changed = true
		p.X = 0
	case upHeld:
		p.Y--
	case downHeld:
		p.Y++
	case leftHeld:
		p.X--
	case rightHeld:
		p.X++
	}

	if p.X != oldX || p.Y != oldY {
		changed = true
	}

	if p.X > 1 {
		p.X = 1
	} else if p.X < -1 {
		p.X = -1
	}
	if p.Y > 1 {
		p.Y = 1
	} else if p.Y < -1 {
		p.Y = -1
	}

	return p, changed
}

func (ce *controlEntity) setAnimation() {
	var (
		upReleased    = engo.Input.Button(upButton).JustReleased()
		downReleased  = engo.Input.Button(downButton).JustReleased()
		leftReleased  = engo.Input.Button(leftButton).JustReleased()
		rightReleased = engo.Input.Button(rightButton).JustReleased()
		anyReleased   = upReleased || downReleased || leftReleased || rightReleased

		upHeld    = engo.Input.Button(upButton).Down()
		downHeld  = engo.Input.Button(downButton).Down()
		leftHeld  = engo.Input.Button(leftButton).Down()
		rightHeld = engo.Input.Button(rightButton).Down()

		upPressed    = engo.Input.Button(upButton).JustPressed()
		downPressed  = engo.Input.Button(downButton).JustPressed()
		leftPressed  = engo.Input.Button(leftButton).JustPressed()
		rightPressed = engo.Input.Button(rightButton).JustPressed()
	)

	switch {
	case upReleased:
		ce.SelectAnimationByAction(StopUpAnimation)
	case downReleased:
		ce.SelectAnimationByAction(StopDownAnimation)
	case leftReleased:
		ce.SelectAnimationByAction(StopLeftAnimation)
	case rightReleased:
		ce.SelectAnimationByAction(StopRightAnimation)
	}

	if anyReleased {
		switch {
		case upHeld:
			ce.SelectAnimationByAction(WalkUpAnimation)
		case downHeld:
			ce.SelectAnimationByAction(WalkDownAnimation)
		case leftHeld:
			ce.SelectAnimationByAction(WalkLeftAnimation)
		case rightHeld:
			ce.SelectAnimationByAction(WalkRightAnimation)
		}
	}

	switch {
	case upPressed:
		ce.SelectAnimationByAction(WalkUpAnimation)
	case downPressed:
		ce.SelectAnimationByAction(WalkDownAnimation)
	case leftPressed:
		ce.SelectAnimationByAction(WalkLeftAnimation)
	case rightPressed:
		ce.SelectAnimationByAction(WalkRightAnimation)
	}
}

// ControlSystem to handle player input for character.
type ControlSystem struct {
	entities []controlEntity
}

// Add an entity to the ControlSystem.
func (cs *ControlSystem) Add(
	b *ecs.BasicEntity,
	a *common.AnimationComponent,
	s *common.SpaceComponent,
	c *ControlComponent,
) {
	cs.entities = append(cs.entities, controlEntity{b, a, s, c})
}

// Remove an entity from the ControlSystem.
func (cs *ControlSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range cs.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		cs.entities = append(cs.entities[:del], cs.entities[del+1:]...)
	}
}

// Update the ControlSystem this frame.
func (cs *ControlSystem) Update(dt float32) {
	for _, e := range cs.entities {
		e.setAnimation()
		if vector, changed := e.speed(); changed {
			vector, _ = vector.Normalize()
			vector.MultiplyScalar(dt * speedScale)
			engo.Mailbox.Dispatch(SpeedMessage{e.BasicEntity, vector})
		}
	}
}

// SpeedMessageType is the unique type identifier for SpeedMessage.
const SpeedMessageType string = "SpeedMessage"

// SpeedMessage updates the speed of an object within the system.
type SpeedMessage struct {
	*ecs.BasicEntity
	engo.Point
}

// Type returns a unique string representation for SpeedMessage.
func (SpeedMessage) Type() string {
	return SpeedMessageType
}

// SpeedComponent stores speed for an entity.
type SpeedComponent struct {
	engo.Point
}

type speedEntity struct {
	*ecs.BasicEntity
	*common.SpaceComponent
	*SpeedComponent
}

// SpeedSystem controls the speed of moving entities.
type SpeedSystem struct {
	entities []speedEntity
}

// New initialises SpeedSystem when it's added to the world.
func (ss *SpeedSystem) New(*ecs.World) {
	engo.Mailbox.Listen(SpeedMessageType, func(msg engo.Message) {
		speed, ok := msg.(SpeedMessage)
		if !ok {
			return
		}

		for _, e := range ss.entities {
			if e.ID() != speed.ID() {
				continue
			}
			e.Point = speed.Point
		}
	})
}

// Add an entity to the SpeedSystem.
func (ss *SpeedSystem) Add(
	basic *ecs.BasicEntity,
	space *common.SpaceComponent,
	speed *SpeedComponent,
) {
	ss.entities = append(ss.entities, speedEntity{basic, space, speed})
}

// Remove an entity from the SpeedSystem.
func (ss *SpeedSystem) Remove(b ecs.BasicEntity) {
	var del int = -1
	for i, e := range ss.entities {
		if e.ID() == b.ID() {
			del = i
			break
		}
	}
	if del >= 0 {
		ss.entities = append(ss.entities[:del], ss.entities[del+1:]...)
	}
}

// Update the SpeedSystem this frame.
func (ss *SpeedSystem) Update(dt float32) {
	speedX := levelWidth * dt
	speedY := levelHeight * dt
	for _, e := range ss.entities {
		e.Position.X = e.Position.X + speedX*e.SpeedComponent.X
		e.Position.Y = e.Position.Y + speedY*e.SpeedComponent.Y

		// limit to map borders
		var limitX float32 = levelWidth - e.Width
		var limitY float32 = levelHeight - e.Height
		switch {
		case e.Position.X < 0:
			e.Position.X = 0
		case e.Position.X > limitX:
			e.Position.X = limitX
		case e.Position.Y < 0:
			e.Position.Y = 0
		case e.Position.Y > limitY:
			e.Position.Y = limitY
		}
	}
}
