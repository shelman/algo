package main

import (
	"fmt"

	"github.com/shelman/heap/impl"
)

func main() {
	heap := impl.NewMinHeap(3)
	heap.Add(3)
	heap.Add(5)
	heap.Add(1)
	heap.Add(4)
	heap.Add(6)
	for !heap.Empty() {
		fmt.Printf("%v\n", heap.RemoveMin())
	}
}
