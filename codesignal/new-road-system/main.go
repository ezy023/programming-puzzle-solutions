package main

func newRoadSystem(roadRegister [][]bool) bool {
	adjacency := make([][]int, len(roadRegister))
	for i := 0; i < len(roadRegister); i++ {
		for j := 0; j < len(roadRegister[i]); j++ {
			if roadRegister[i][j] == true {
				adjI := adjacency[i]
				if len(adjI) == 0 {
					buf := make([]int, 1)
					buf[0] = j
					adjacency[i] = buf
				} else {
					adjI = append(adjI, j)
					adjacency[i] = adjI
				}
			}
		}
	}

	outgoing := make([]int, len(roadRegister))
	for _, in := range adjacency {
		for _, i := range in {
			outgoing[i]++
		}
	}

	for i := 0; i < len(adjacency); i++ {
		if outgoing[i] != len(adjacency[i]) {
			return false
		}
	}

	return true
}
