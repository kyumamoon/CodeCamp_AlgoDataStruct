package main

import (
	"fmt"
	"math"
	"sort"
)

// binarySearch returns the index and true or false whether the target is found in the list.
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

// recursiveBinarySearch returns true or false whether target is found in the list using recursive.
func recursiveBinarySearch(list []int, target int) bool {
	if len(list) == 0 {
		return false
	} else {
		midpoint := math.Floor(float64(len(list)) / 2)

		if list[int(midpoint)] == target {
			return true
		} else if list[int(midpoint)] < target {
			return recursiveBinarySearch(list[int(midpoint)+1:], target)
		} else {
			return recursiveBinarySearch(list[:int(midpoint)], target)
		}
	}

}

func main() {
	list := []int{1, 7, 2, 4, 3, 8, 6, 5}
	// Sorted: 1,2,3,4,5,6,7,8
	sort.Ints(list)
	/*x, found := binarySearch(list, 10)
	if found {
		fmt.Println("INDEX:", x)
	} else {
		fmt.Println("NO VALUES FOUND.")
	}*/

	y := recursiveBinarySearch(list, 11)
	fmt.Println("Target Found:", y)
}
