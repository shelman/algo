package main

import "fmt"

func main() {
	fmt.Printf("%#v\n", makeChange([]int{1, 2, 5}, 56))
}

func makeChange(coins []int, total int) int {
	memoization := map[int]int{}
	return _makeChange(coins, total, memoization)
}

func _makeChange(coins []int, total int, memoization map[int]int) int {
	if existingAnswer, ok := memoization[total]; ok {
		return existingAnswer
	}

	if total == 0 {
		return 0
	}

	if total < 0 {
		return -1
	}

	fewestCoinsNeeded := -1
	for _, coin := range coins {
		fewestCoinsSubtotal := _makeChange(coins, total - coin, memoization)
		if fewestCoinsSubtotal != -1 && (fewestCoinsSubtotal < fewestCoinsNeeded || fewestCoinsNeeded == -1) {
			fewestCoinsNeeded = fewestCoinsSubtotal + 1
		}
	}

	memoization[total] = fewestCoinsNeeded
	return fewestCoinsNeeded
}
