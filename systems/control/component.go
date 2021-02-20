package control

// ControlComponent stores control input for player entity.
type ControlComponent struct {
	Vertical, Horizontal string
	Enabled              bool
}

// GetControlComponent provides type safe access to ControlComponent.
func (c *ControlComponent) GetControlComponent() *ControlComponent {
	return c
}
