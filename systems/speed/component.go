package speed

import "github.com/EngoEngine/engo"

// SpeedComponent stores speed for an entity.
type SpeedComponent struct {
	engo.Point
}

// GetSpeedComponent provides type safe access to SpeedComponent.
func (s *SpeedComponent) GetSpeedComponent() *SpeedComponent {
	return s
}
