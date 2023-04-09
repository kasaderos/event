package event

import "testing"

func TestManager_Initiate(t *testing.T) {
	type args struct {
		e Event
	}

	// 0 -> 1 -> 2
	// 3 -> 1
	// run from 0
	count := 0
	manager := NewManager()

	v0 := Event{
		"0",
		func() { count++ },
	}
	v1 := Event{
		"1",
		nil,
	}
	v2 := Event{
		"2",
		func() { count++ },
	}
	v3 := Event{
		"3",
		func() { count++ },
	}

	manager.Bind(v0, v1)
	manager.Bind(v1, v2)
	manager.Bind(v3, v1)

	tests := []struct {
		name string
		s    *Manager
		args args
	}{
		{"event chain", manager, args{v0}},
		{"one event", manager, args{v3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.s.Initiate(tt.args.e)
			if tt.name == "event chain" && count != 2 {
				t.Errorf("expected 2 actual %d\n", count)
			}
			if tt.name == "one event" && count != 4 {
				t.Errorf("expected 4 actual %d\n", count)
			}
		})
	}

}
