package main

func phoneCall(min1, min2_10, min11, s int) int {
	reduce := func(length, cost int) int {
		if length == 0 {
			return s / cost
		} else if s >= (length * cost) {
			s = s - (length * cost)
			return length
		} else {
			minutes := s / cost
			s = s - (minutes * cost)
			return minutes
		}
	}

	sum := 0
	sum += reduce(1, min1)
	sum += reduce(9, min2_10)
	if sum >= 10 {
		sum += reduce(0, min11)
	}

	return sum
}
