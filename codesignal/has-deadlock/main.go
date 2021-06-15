package main

func hasDeadlock(connections [][]int) bool {
	contains := func(col []int, i int) bool {
		for _, item := range col {
			if i == item {
				return true
			}
		}
		return false
	}

	var check func(node int, visited []int) bool

	check = func(node int, visited []int) bool {
		if contains(visited, node) {
			return true
		}
		visited = append(visited, node)
		for _, n := range connections[node] {
			if check(n, visited) {
				return true
			}
		}
		return false
	}

	seen := make([]int, 0)
	for idx, _ := range connections {
		if check(idx, seen) {
			return true
		}
	}

	return false
}
