package main

import "fmt"

func main() {
	fmt.Printf("%v\n", binSearch([]int{1, 3, 4, 5, 8, 10, 25, 34, 55}, 5))
}

func _binSearch(list []int, target int, start int, end int) int {
	if start == end {
		if list[start] == target {
			return start
		}
		return -1
	}

	idx := (start+end)/2
	guess := list[idx]

	if guess == target {
		return idx
	}

	if guess < target {
		return _binSearch(list, target, idx+1, end)
	}

	return _binSearch(list, target, start, idx-1)
}

func binSearch(list []int, target int) int {
	if len(list) == 0 {
		return -1
	}

	return _binSearch(list, target, 0, len(list)-1)
}
