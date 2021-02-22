package character

// ActionType indicates the type of an Action.
type ActionType int

// Actions for NPCs.
const (
	ActStop ActionType = iota
	ActWalk
	ActRun
)

// An Action defines something an NPC does.
type Action struct {
	Type ActionType
}

// Schedule contains an Action list for NPCS to carry out.
type Schedule struct {
	Actions       []Action
	currentAction int
}

// CurrentAction returns the current Action for the Schedule.
func (s Schedule) CurrentAction() *Action {
	if s.currentAction >= len(s.Actions) {
		return nil
	}
	return &s.Actions[s.currentAction]
}
