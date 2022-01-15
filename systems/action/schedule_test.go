package action

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

	t.Run("StepAction", func(t *testing.T) {
		schedule.currentAction = 1
		schedule.currentDuration = 1
		schedule.StepAction()

		wantAction, gotAction := 2, schedule.currentAction
		if wantAction != gotAction {
			t.Errorf("want CurrentAction == %v, got %v", wantAction, gotAction)
		}

		wantDuration, gotDuration := float32(0), schedule.currentDuration
		if wantDuration != gotDuration {
			t.Errorf("want CurrentDuration == %v, got %v", wantDuration, gotDuration)
		}

		schedule.Loop = true
		schedule.currentDuration++
		schedule.StepAction()

		wantAction, gotAction = 0, schedule.currentAction
		if wantAction != gotAction {
			t.Errorf("want CurrentAction == %v, got %v", wantAction, gotAction)
		}

		wantDuration, gotDuration = float32(0), schedule.currentDuration
		if wantDuration != gotDuration {
			t.Errorf("want CurrentDuration == %v, got %v", wantDuration, gotDuration)
		}
	})
}
