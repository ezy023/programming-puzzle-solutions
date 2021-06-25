package main

import "fmt"

func treeDiameter(n int, tree [][]int) int {
	graph := make(map[int][]int)

	for _, t := range tree {
		if _, ok := graph[t[0]]; !ok {
			v := make([]int, 0, 1)
			v = append(v, t[1])
			graph[t[0]] = v
		} else {
			graph[t[0]] = append(graph[t[0]], t[1])
		}
	}

	max := func(col []int) int {
		tmp := 0
		for _, i := range col {
			if i > tmp {
				tmp = i
			}
		}
		return tmp
	}

	maxtwo := func(vals []int) []int {
		maxes := make([]int, 2)
		for _, v := range vals {
			if v > maxes[0] {
				maxes[0], maxes[1] = v, maxes[0]
			} else if v > maxes[1] {
				maxes[1] = v
			}
		}
		return maxes
	}

	var calcHeight func(node int) int
	diameters := make(map[int]int)
	calcHeight = func(node int) int {
		vertices, ok := graph[node]
		if !ok {
			return 0
		}

		heights := make([]int, len(vertices))
		for idx, vert := range vertices {
			heights[idx] = calcHeight(vert)
		}
		fmt.Printf("Heights: %#v\n", heights)

		maxHeight := max(heights) + 1
		two := maxtwo(heights)
		diameters[node] = two[0] + two[1] + 1
		return maxHeight
	}

	for node, _ := range graph {
		calcHeight(node)
	}

	d := make([]int, 0)
	for _, v := range diameters {
		d = append(d, v)
	}

	return max(d)

	// calculatedHeights := make(map[int]int)
	// for node, _ := range graph {
	// 	h := calcHeight(node)
	// 	calculatedHeights[node] = h
	// }

	// var maxDiam func(node int) int

	// maxDiam = func(node int) int {
	// 	if vertices, ok := graph[node]; !ok {
	// 		return 0
	// 	}

	// 	heights := make([]int, len(vertices))
	// 	for idx, vert := range vertices {
	// 		heights
	// 	}
	// }
}

func treeDiameterOne(n int, tree [][]int) int {
	graph := make(map[int][]int)

	for _, t := range tree {
		if _, ok := graph[t[0]]; !ok {
			v := make([]int, 0, 1)
			v = append(v, t[1])
			graph[t[0]] = v
		} else {
			graph[t[0]] = append(graph[t[0]], t[1])
		}
	}

	maxtwo := func(vals []int) []int {
		maxes := make([]int, 2)
		for _, v := range vals {
			if v > maxes[0] {
				maxes[0], maxes[1] = v, maxes[0]
			} else if v > maxes[1] {
				maxes[1] = v
			}
		}
		return maxes
	}

	calculatedMax := make(map[int]int)
	// need to memoize
	var maxlen func(node int) int
	maxlen = func(node int) int {
		if v, ok := calculatedMax[node]; ok {
			return v
		}
		vertices, ok := graph[node]
		if !ok {
			calculatedMax[node] = 0
			return 0
		}
		lengths := make([]int, len(vertices))
		for idx, vert := range vertices {
			lengths[idx] = maxlen(vert)
		}
		topTwo := maxtwo(lengths)
		maxPath := topTwo[0] + topTwo[1] + 1
		calculatedMax[node] = maxPath
		return maxPath
	}

	absMax := 0
	for node, _ := range graph {
		nodeLen := maxlen(node)
		if nodeLen > absMax {
			absMax = nodeLen
		}
	}
	return absMax
}
