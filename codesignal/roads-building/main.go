package main

func roadsBuilding(cities int, roads [][]int) [][]int {
	contains := func(n int, set []int) bool {
		for _, i := range set {
			if i == n {
				return true
			}
		}
		return false
	}
	currentRoads := make(map[int][]int, cities)
	for _, road := range roads {
		col := currentRoads[road[0]]
		if !contains(road[1], col) {
			col = append(col, road[1])
			currentRoads[road[0]] = col
		}

		col = currentRoads[road[1]]
		if !contains(road[0], col) {
			col = append(col, road[0])
			currentRoads[road[1]] = col
		}
	}

	buf := make([][]int, 0, cities)
	for i := 0; i < cities; i++ {
		cityRoads := currentRoads[i]
		if len(cityRoads) < cities {
			for j := i + 1; j < cities; j++ {
				if !contains(j, cityRoads) {
					road := make([]int, 2)
					road[0], road[1] = i, j
					buf = append(buf, road)
					currentRoads[i] = append(currentRoads[i], j)
					currentRoads[j] = append(currentRoads[j], i)
				}
			}
		}
	}

	return buf
}
