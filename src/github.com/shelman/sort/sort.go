package main

import "fmt"

func main() {
	input := []int{3, 5, 2, -1, 0, 3, 9, -2}

	fmt.Printf("%v\n", quicksort(input))
}

func quicksort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	highs := []int{}
	mids := []int{}
	lows := []int{}

	pivot := input[len(input)/2]
	for _, i := range input {
		if i > pivot {
			highs = append(highs, i)
		} else if i == pivot {
			mids = append(mids, i)
		} else {
			lows = append(lows, i)
		}
	}
	return append(quicksort(lows), append(mids, quicksort(highs)...)...)
}
