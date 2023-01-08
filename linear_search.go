package main

import "fmt"

// linear search searches a list for the target and returns the index in the list if found
func linear_search(list []int, target int) (int, bool) {
	for i := 0; i < len(list); i++ {
		if list[i] == target {
			return i, true
		}
	}
	return 0, false
}

func main() {
	list := []int{1, 412, 41, 561, 612, 513, 51, 421, 1, 31, 31, 51, 51, 351, 51, 51, 2421, 41, 4124}
	x, found := linear_search(list, 22)
	if found {
		fmt.Println("INDEX:", x)
	} else {
		fmt.Println("NO VALUES FOUND.")
	}
}
