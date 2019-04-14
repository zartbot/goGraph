package gograph

import "math"

const (
	INFINITY    = math.MaxInt64 //int64(^uint64(0) >> 1)
	UNINITVALUE = -1
)

//Vertex is a node in Graph
type Vertex int

//Graph structure include Label/Vertex mapping and edges
type Graph struct {
	VertexMap map[string]Vertex
	LabelMap  map[Vertex]string
	Edges     map[Vertex]map[Vertex]int64
	VertexNum int
}

//NewGraph is used for create graph
func NewGraph() *Graph {
	g := &Graph{
		VertexMap: make(map[string]Vertex),
		LabelMap:  make(map[Vertex]string),
		Edges:     make(map[Vertex]map[Vertex]int64),
		VertexNum: 0,
	}
	return g
}

//GraphIntf is graph interface for algorithm fetch data
type GraphIntf interface {
	VertexList() []Vertex
	Neighbors(v Vertex) []Vertex
	Weight(u, v Vertex) int64
}

//VertexList is used to generate all vertex in a list
func (g *Graph) VertexList() (vlist []Vertex) {
	for _, v := range g.VertexMap {
		vlist = append(vlist, v)
	}
	return vlist
}

//Neighbors will be used for generate adjacency map
func (g *Graph) Neighbors(u Vertex) (vlist []Vertex) {
	for v := range g.Edges[u] {
		vlist = append(vlist, v)
	}
	return vlist
}

//Weight will be used for generate adjacency map
func (g *Graph) Weight(u, v Vertex) int64 {
	return g.Edges[u][v]
}
