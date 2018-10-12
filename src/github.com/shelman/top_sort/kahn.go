package main

import (
	"fmt"

	"github.com/shelman/graph"
	"github.com/shelman/queue"
)

func main() {
	dag := &graph.Graph{}
	dag.AddNode("a")
	dag.AddNode("b")
	dag.AddNode("c")
	dag.AddNode("d")
	dag.AddNode("e")

	dag.AddEdge("a", "b")
	dag.AddEdge("a", "c")
	dag.AddEdge("b", "c")
	dag.AddEdge("d", "b")
	dag.AddEdge("d", "e")

	fmt.Printf("%v\n", topSort(dag))
}

func topSort(dag *graph.Graph) []string {
	result := []string{}

	sources := dag.Sources()
	sourceQ := queue.NewStrQueue(sources)
	for !sourceQ.Empty() {
		source := sourceQ.Pop()
		result = append(result, source)

		for _, sink := range dag.EdgesFrom(source) {
			dag.RemoveEdge(source, sink)
			if len(dag.EdgesTo(sink)) == 0 {
				sourceQ.Push(sink)
			}
		}
	}

	if dag.NumEdges() > 0 {
		panic("Cycle in the graph")
	}

	return result
}


