package magic

import (
	"github.com/EngoEngine/ecs"
)

// Message updates an object within the system.
type Message struct {
	*ecs.BasicEntity
}

// Type returns a unique string representation for Message.
func (Message) Type() string {
	return MessageType
}
