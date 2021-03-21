package main

func addBorder(picture []string) []string {
	addBorderRow := func(w int) string {
		collector := make([]byte, 0, w)
		for i := 0; i < w; i++ {
			collector = append(collector, '*')
		}
		return string(collector)
	}

	width := len(picture[0]) + 2
	collector := make([]string, 0, len(picture)+2)
	collector = append(collector, addBorderRow(width))
	for _, row := range picture {
		rowCollect := make([]byte, 0, width)
		rowCollect = append(rowCollect, '*')
		for j := 0; j < len(row); j++ {
			rowCollect = append(rowCollect, row[j])
		}
		rowCollect = append(rowCollect, '*')
		collector = append(collector, string(rowCollect))
	}
	collector = append(collector, addBorderRow(width))
	return collector
}
