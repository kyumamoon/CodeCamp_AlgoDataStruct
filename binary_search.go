package main

import (
	"fmt"
	"math"
	"sort"
)

func binarySearch(list []int, target int) (int, bool) {
	first := 0
	last := len(list) - 1

	for first <= last {
		midpoint := math.Floor(float64(first+last) / 2)

		if list[int(midpoint)] == target {
			return int(midpoint), true
		} else if list[int(midpoint)] < target {
			first = int(midpoint) + 1
		} else {
			last = int(midpoint) - 1
		}

	}

	return 0, false

}

func main() {
	list := []int{1, 7, 2, 4, 3, 8, 6, 5, 9, 10}
	// Sorted: 1,2,3,4,5,6,7,8,9,10
	sort.Ints(list)
	x, found := binarySearch(list, 10)
	if found {
		fmt.Println("INDEX:", x)
	} else {
		fmt.Println("NO VALUES FOUND.")
	}
}
