package character

// CharacterComponent marks entities for use with a CharacterSystem.
type CharacterComponent struct {
	Schedule
}

// GetCharacterComponent provides type safe access to CharacterComponent.
func (pc *CharacterComponent) GetCharacterComponent() *CharacterComponent {
	return pc
}
