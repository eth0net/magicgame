package character

import (
	"reflect"
	"testing"
)

func TestSchedule(t *testing.T) {
	actions := []Action{{Type: ActStop}}

	schedule := Schedule{
		Actions: actions,
	}

	t.Run("CurrentAction", func(t *testing.T) {
		want, got := &actions[0], schedule.CurrentAction()
		if !reflect.DeepEqual(want, got) {
			t.Errorf("want CurrentAction == %v, got %v", want, got)
		}

		schedule.currentAction++
		want, got = nil, schedule.CurrentAction()
		if want != got {
			t.Errorf("want CurrentAction == %v, got %v", want, got)
		}
	})
}
