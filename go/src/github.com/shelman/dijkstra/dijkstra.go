package main

import (
	"fmt"
	"math"

	"github.com/shelman/graph"
)

func main() {
	fmt.Printf("Hello\n")

	g := &graph.Graph{}
	g.AddNode("a")
	g.AddNode("b")
	g.AddNode("c")
	g.AddNode("d")
	g.AddNode("e")

	g.AddWeightedEdge("a", "b", 1)
	g.AddWeightedEdge("a", "c", 1)
	g.AddWeightedEdge("b", "d", 1)
	g.AddWeightedEdge("d", "e", 1)
	g.AddWeightedEdge("c", "e", 1)

	runDijkstra(g)
}

func getNeighbors(g *graph.Graph, node string, unvisitedNodes map[string]bool) []string {
	from := g.EdgesFrom(node)

	neighbors := []string{}
	for _, neighbor := range from {
		if !unvisitedNodes[neighbor] {
			continue
		}
		neighbors = append(neighbors, neighbor)
	}
	return neighbors
}

func runDijkstra(g *graph.Graph) {
	nodes := g.Nodes()

	unvisitedNodes := map[string]bool{}
	distances := map[string]int{}
	for _, node := range nodes {
		unvisitedNodes[node] = true
		distances[node] = 10000
	}

	currentNode := nodes[0]
	distances[currentNode] = 0

	for true {
		unvisitedNodes[currentNode] = false
		neighbors := getNeighbors(g, currentNode, unvisitedNodes)

		for _, neighbor := range neighbors {
			distances[neighbor] = int(math.Min(float64(distances[neighbor]),
				float64(distances[currentNode]+g.Weight(currentNode, neighbor))))
		}

		smallestDist := 10000
		smallestDistNode := ""
		for node, needsVisit := range unvisitedNodes {
			if needsVisit == false {
				continue
			}
			if distances[node] < smallestDist {
				smallestDist = distances[node]
				smallestDistNode = node
			}
		}

		if smallestDist == 10000 {
			break
		}

		currentNode = smallestDistNode
	}

	fmt.Printf("%#v\n", distances)
}
