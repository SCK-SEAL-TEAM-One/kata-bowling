package bowling_test

import (
	. "bowling"
	"testing"
)

func Test_Bowling_12_Rolls_All_Strike_Should_Be_300_Points(t *testing.T) {
	expectedPoints := 300
	frames := []Frame{
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X"},
		},
		Frame{
			Rolls: []string{"X", "X", "X"},
		},
	}

	actualPoints := CalculatePoint(frames)

	if expectedPoints != actualPoints {
		t.Errorf("expect %d but it got %d", expectedPoints, actualPoints)
	}
}
func Test_Bowling_20_Rolls_Should_Be_90_Points(t *testing.T) {
	expectedPoints := 90
	frames := []Frame{
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
		Frame{
			Rolls: []string{"9", "-"},
		},
	}

	actualPoints := CalculatePoint(frames)

	if expectedPoints != actualPoints {
		t.Errorf("expect %d but it got %d", expectedPoints, actualPoints)
	}
}
