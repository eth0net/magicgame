package magic

// Component handles magic for an entity.
type Component struct{}

// GetComponent provides type safe access to Component.
func (c *Component) GetComponent() *Component {
	return c
}
