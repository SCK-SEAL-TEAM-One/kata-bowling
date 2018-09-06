package bowling_test

import (
	. "bowling"
	"testing"
)

func Test_Calculate_12_Rolls_All_Strike_Should_Be_300_Point(t *testing.T) {
	expectedPoints := 300

	inputPlay := []int{30, 30, 30, 30, 30, 30, 30, 30, 30, 30}

	actualResult := Calculate(inputPlay)
	if expectedPoints != actualResult {
		t.Errorf("expected %d but got it %d", expectedPoints, actualResult)
	}
}
