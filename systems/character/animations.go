package character

import "github.com/EngoEngine/engo/common"

// Animation names for Characters to implement,
// used by CharacterSystem to trigger Character Animations.
const (
	AnimationMoveUp    string = "MoveUp"
	AnimationMoveDown  string = "MoveDown"
	AnimationMoveLeft  string = "MoveLeft"
	AnimationMoveRight string = "MoveRight"
	AnimationStopUp    string = "StopUp"
	AnimationStopDown  string = "StopDown"
	AnimationStopLeft  string = "StopLeft"
	AnimationStopRight string = "StopRight"
)

// CharacterAnimations delcares default Character Animations.
var CharacterAnimations = []*common.Animation{
	{Name: AnimationMoveUp, Frames: []int{9, 10, 11}, Loop: true},
	{Name: AnimationMoveDown, Frames: []int{0, 1, 2}, Loop: true},
	{Name: AnimationMoveLeft, Frames: []int{3, 4, 5}, Loop: true},
	{Name: AnimationMoveRight, Frames: []int{6, 7, 8}, Loop: true},
	{Name: AnimationStopUp, Frames: []int{10}, Loop: true},
	{Name: AnimationStopDown, Frames: []int{1}, Loop: true},
	{Name: AnimationStopLeft, Frames: []int{4}, Loop: true},
	{Name: AnimationStopRight, Frames: []int{7}, Loop: true},
}
