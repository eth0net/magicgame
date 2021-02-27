package character

// A Schedule contains a list of Actions for Characters to perform.
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
