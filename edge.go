package gograph

//AddEdge is used add edge for graph
func (g *Graph) AddEdge(u, v string, metric int64) {
	if _, ok := g.Edges[g.VertexMap[u]]; !ok {
		g.Edges[g.VertexMap[u]] = make(map[Vertex]int64)
	}
	g.Edges[g.VertexMap[u]][g.VertexMap[v]] = metric
}
