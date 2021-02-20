package game

import (
	"testing"
)

var scene = Scene{}

func TestGameScene_Type(t *testing.T) {
	var want, got string = SceneType, scene.Type()
	if want != got {
		t.Errorf("want type == %v, got %v", want, got)
	}
}
