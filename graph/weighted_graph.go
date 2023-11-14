package graph

type WeightedGraph struct {
	distances map[Pair]int
	nodes     []Node
}

type Node struct {
	name  string
	value int
}

type Pair struct {
	n1 Node
	n2 Node
}
