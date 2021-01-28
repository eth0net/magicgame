package scenes

import (
	"testing"
)

var scene GameScene = GameScene{}

func TestGameScene_Type(t *testing.T) {
	var want, got string = GameSceneType, scene.Type()
	if want != got {
		t.Errorf("want type == %v, got %v", want, got)
	}
}
