package magic

import (
	"github.com/EngoEngine/ecs"
	"github.com/EngoEngine/engo"
	"github.com/EngoEngine/engo/common"
	"github.com/eth0net/magicgame/entities/fireball"
)

type System struct {
	world    *ecs.World
	entities map[uint64]entity
}

func (s *System) New(world *ecs.World) {
	s.world = world

	engo.Mailbox.Listen(MessageType, func(msg engo.Message) {
		m := msg.(Message)
		e, ok := s.entities[m.ID()]
		if !ok {
			return
		}
		// TODO: cast spell
		f := fireball.New(e.Position, engo.Point{X: 0.02})
		s.world.AddEntity(f)
	})

	engo.Mailbox.Listen("CollisionMessage", func(msg engo.Message) {
		m := msg.(common.CollisionMessage)
		if _, ok := s.entities[m.Entity.ID()]; !ok {
			return
		}
		// TODO: handle collision for spells
	})
}

// Add an entity to the System.
func (s *System) Add(
	basic *ecs.BasicEntity,
	space *common.SpaceComponent,
	magic *Component,
) {
	if s.entities == nil {
		s.entities = map[uint64]entity{}
	}
	s.entities[basic.ID()] = entity{
		basic,
		space,
		magic,
	}
}

// AddByInterface adds entities to the System via Able interface.
func (s *System) AddByInterface(i ecs.Identifier) {
	e := i.(Able)
	s.Add(
		e.GetBasicEntity(),
		e.GetSpaceComponent(),
		e.GetComponent(),
	)
}

func (s *System) Remove(e ecs.BasicEntity) {
	if _, ok := s.entities[e.ID()]; !ok {
		return
	}
	delete(s.entities, e.ID())
}

func (s *System) Update(dt float32) {
	// TODO: update any spells - maybe TTL expiry etc
}
