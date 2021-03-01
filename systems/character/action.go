package character

import (
	"github.com/EngoEngine/engo"
)

// ActionType indicates the type of an Action.
type ActionType int

// ActionTypes available for use in Character Schedules.
const (
	// ActStop stops Character movement and turns the
	// Character to face the direction indicated by Target.
	ActStop ActionType = iota

	// ActTurn turns the Character to face the direction
	// indicated by Target and updates the SpeedComponent.
	ActTurn

	// ActWalk makes the Character walk in the direction
	// indicated by Target until Duration has passed.
	ActWalk

	// ActRun makes the Character run in the direction
	// indicated by Target until Duration has passed.
	ActRun

	// ActTurnTo turns the Character to face the location
	// indicated by Target and updates the SpeedComponent.
	ActTurnTo

	// ActWalkTo makes the Character walk to the location
	// indicated by Target, completing upon arrival.
	ActWalkTo

	// ActRunTo makes the Character run to the location
	// indicated by Target, completing upon arrival.
	ActRunTo

	// ActTeleportTo instantly teleports the Character to the
	// location indicated by Target, completing upon arrival.
	ActTeleportTo

	// Other ideas:
	// FollowPath
	// FollowCharacter
	// FollowSpaceComponent
	// Interact
	// Attack
	// Defend
	// Magic
	// Jump
	// Anything else
)

// An Action defines a single act for a Character.
// Actions are used to create Character Schedules.
type Action struct {
	// Type determines what to do to the Character
	// and how to interpret the other Action fields.
	Type ActionType

	// Point is the relevant engo.Point for the action.
	// ActionTypes can use Point in different ways,
	// refer to the ActionType for more information.
	//
	// Examples:
	//  - ActWalk uses Point as a direction.
	//  - ActWalkTo uses Point as a destination.
	//
	// When Point is a direction, the Action will be complete
	// after one system update, unless Duration is greater than 0.
	//
	// When Point is a destination, the Action will be complete
	// when the Point is reached and Duration has no effect.
	engo.Point

	// Duration sets how long the Action should run for,
	// if Duration is 0 the Action runs until complete.
	// Note that Duration does not affect all ActionTypes.
	Duration float32
}
