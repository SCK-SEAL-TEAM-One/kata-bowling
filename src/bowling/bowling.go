package bowling

type Frame struct {
	Rolls []string
}

func CalculatePoint(frames []Frame) int {
	sum := 0
	for _, frame := range frames {
		if frame.Rolls[0] == "X" {
			sum += 30
		}
	}
	return sum
}
