package action

// ActionComponent marks entities for use with a ActionSystem.
type ActionComponent struct {
	Schedule
}

// GetActionComponent provides type safe access to ActionComponent.
func (ac *ActionComponent) GetActionComponent() *ActionComponent {
	return ac
}
