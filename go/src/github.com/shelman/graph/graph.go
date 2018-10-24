package graph

type edge struct {
	from string
	to string
	weight int
}

type Graph struct {
	nodes []string
	edges []edge
}

func (g *Graph) AddNode(nodeId string) {
	g.nodes = append(g.nodes, nodeId)
}

func (g *Graph) AddEdge(source string, sink string) {
	g.AddWeightedEdge(source, sink, 1)
}

func (g *Graph) AddWeightedEdge(source string, sink string, weight int) {
	g.edges = append(g.edges, edge{source, sink, weight})
}

func (g *Graph) EdgesFrom(source string) []string {
	result := []string{}
	for _, edge := range g.edges {
		if edge.from == source {
			result = append(result, edge.to)
		}
	}
	return result
}

func (g *Graph) EdgesTo(sink string) []string {
	result := []string{}
	for _, edge := range g.edges {
		if edge.to == sink {
			result = append(result, edge.from)
		}
	}
	return result
}

func (g *Graph) NumEdges() int {
	return len(g.edges)
}

func (g *Graph) Nodes() []string {
	return g.nodes
}

func (g *Graph) RemoveEdge(source string, sink string) {
	for i, edge := range g.edges {
		if edge.from == source && edge.to == sink {
			g.edges = append(g.edges[:i], g.edges[i+1:]...)
		}
	}
}

func (g *Graph) Sources() []string {
	isSource := map[string]bool{}
	for _, n := range g.nodes {
		isSource[n] = true
	}
	for _, e := range g.edges {
		isSource[e.to] = false
	}

	sources := []string{}
	for n, isSource := range isSource {
		if isSource {
			sources = append(sources, n)
		}
	}
	return sources
}

func (g *Graph) Weight(source string, sink string) int {
	for _, edge := range g.edges {
		if edge.from == source && edge.to == sink {
			return edge.weight
		}
	}
	panic("Edge not found")
}
