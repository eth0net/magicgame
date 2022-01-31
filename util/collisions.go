package util

import "github.com/EngoEngine/engo/common"

const (
	CollisionWorld common.CollisionGroup = 1 << iota
	CollisionPlayer
	CollisionEntity
)
