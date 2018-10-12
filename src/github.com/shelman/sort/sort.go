package main

import "fmt"

func main() {
	input := []int{3, 5, 2, -1, 0, 3, 9, -2}

	fmt.Printf("%v\n", quicksort(input))
	fmt.Printf("%v\n", mergesort(input))
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

func merge(list1 []int, list2 []int) []int {
	res := []int{}
	p1 := 0
	p2 := 0
	for i := 0; i < len(list1) + len(list2); i++ {
		if p1 < len(list1) && (p2 >= len(list2) || list1[p1] < list2[p2]) {
			res = append(res, list1[p1])
			p1++
		} else {
			res = append(res, list2[p2])
			p2++
		}
	}
	return res
}

func mergesort(input []int) []int {
	if len(input) == 0 || len(input) == 1 {
		return input
	}

	return merge(mergesort(input[:len(input)/2]), mergesort(input[len(input)/2:]))
}