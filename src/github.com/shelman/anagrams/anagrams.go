package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	D := []string{"a", "baca", "caba", "cabad", "caca", "bd"}
	W := "abca"
	fmt.Printf("%v\n", anagrams(D, W))
}

func anagrams(D []string, W string) []string {
	W_len := len(W)
	W_sorted := sortString(W)
	result := []string{}
	for _, D_word := range D {
		if len(D_word) != W_len {
			continue
		}
		if sortString(D_word) == W_sorted {
			result = append(result, D_word)
		}
	}
	return result
}

func sortString(s string) string {
	sslice := strings.Split(s, "")
	sort.Strings(sslice)
	return strings.Join(sslice, "")
}
