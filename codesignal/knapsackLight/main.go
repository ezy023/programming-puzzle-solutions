package main

func knapsackLight(value1, weight1, value2, weight2, maxW int) int {
	if weight1 > maxW && weight2 > maxW {
		return 0
	}
	if weight1+weight2 <= maxW {
		return value1 + value2
	}
	if value1 > value2 {
		if weight1 <= maxW {
			return value1
		}
		return value2
	}

	if weight2 <= maxW {
		return value2
	}
	return value1

}
