package action

import (
	"github.com/EngoEngine/engo"
)

// ActionType indicates the type of an Action.
type ActionType int

// ActionTypes available for use in Schedules.
const (
	// ActStop stops movement, preserving the direction.
	ActStop ActionType = iota

	// ActTurn turns the Entity to face the direction
	// indicated by Point and stops movement.
	ActTurn

	// ActWalk makes the Entity walk in the direction
	// indicated by Point until Duration has passed.
	ActWalk

	// ActRun makes the Entity run in the direction
	// indicated by Point until Duration has passed.
	ActRun

	// ActTeleport teleports the Entity in the direction
	// indicated by Point until Duration has passed.
	ActTeleport

	// ActTurnTo turns the Entity to face the location
	// indicated by Point and updates the SpeedComponent.
	ActTurnTo

	// ActWalkTo walks the Entity to the location indicated by Point.
	ActWalkTo

	// ActRunTo runs the Entity to the location indicated by Point.
	ActRunTo

	// ActTeleportTo teleports the Entity to the location indicated by Point.
	ActTeleportTo

	// Other ideas:
	// FollowPath
	// FollowEntity
	// FollowSpaceComponent
	// Interact
	// Attack
	// Defend
	// Magic
	// Jump
	// Anything else
)

// An Action defines a single act for a Entity.
// Actions are used to create Schedules.
type Action struct {
	// Type determines what to do to the Entity
	// and how to interpret the other Action fields.
	Type ActionType

	// Point is the relevant engo.Point for the action.
	// ActionTypes can use Point in different ways,
	// refer to the ActionType for more information.
	//
	// Examples:
	//  - ActWalk uses Point as a direction.
	//  - ActWalkTo uses Point as a location.
	//
	// When Point is a direction, the Action will be complete
	// after the Duration has passed. If Duration == 0, the
	// Action will be complete after one system update.
	//
	// When Point is a location, the Action will be complete
	// after the Point is reached and the Duration has passed.
	engo.Point

	// Duration sets how long the Action should run for,
	// if Duration is 0 the Action runs until complete.
	// Note that Duration may not affect all ActionTypes.
	Duration float32
}
