package action

// A Schedule contains a list of Actions to perform.
type Schedule struct {
	Actions         []Action
	Loop            bool
	currentAction   int
	currentDuration float32
}

// CurrentAction returns the current Action for the Schedule.
func (s *Schedule) CurrentAction() *Action {
	if s.currentAction >= len(s.Actions) {
		return nil
	}
	return &s.Actions[s.currentAction]
}

// StepAction steps the Schedule to the next Action.
// If the current Action is the final Action in the
// Schedule, the Schedule starts from the beginning.
func (s *Schedule) StepAction() {
	s.currentAction++
	if s.Loop && s.currentAction >= len(s.Actions) {
		s.currentAction = 0
	}
	s.currentDuration = 0
}
