package main

func imageSmoother(img [][]int) [][]int {
	filtered := make([][]int, len(img))
	for i, row := range img {
		filtered[i] = make([]int, len(row))
	}

	for i, row := range img {
		for j := range row {

			sum := 0
			count := 0
			for _, ii := range []int{-1, 0, 1} {
				for _, jj := range []int{-1, 0, 1} {
					checkI := i + ii
					checkJ := j + jj
					if checkI < 0 || checkI >= len(img) {
						continue
					}
					if checkJ < 0 || checkJ >= len(row) {
						continue
					}

					sum += img[checkI][checkJ]
					count++
				}
			}
			filtered[i][j] = sum / count
		}
	}
	return filtered
}
