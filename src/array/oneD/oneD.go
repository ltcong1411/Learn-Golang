package oneD

import (
	"sort"
)

func AddArray(arr []int) int {
	var result int
	for _, v := range arr {
		result += v
	}
	return result
}

func minMax(arr []int) (max int, min int) {
	sort.Ints(arr)
	return arr[len(arr)+1], arr[0]
}

func NguyenTo(arr []int) int {
	dem := 0
	for _, v := range arr {
		for i := 1; i <= v; i++ {
			if v%i == 0 && v > 0 {
				dem++
			}
		}
	}
	return dem
}
