package bowling

type Frame struct {
	Rolls []string
}

func Calculate(point []int) int {
	sum := 0
	for _, num := range point {
		sum += num
	}
	return sum
}
