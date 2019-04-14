package gograph

import "container/heap"

func Dijkstra(g *Graph, source Vertex) (dist map[Vertex]int64, prev map[Vertex]Vertex) {
	dist = make(map[Vertex]int64)
	prev = make(map[Vertex]Vertex)
	sid := source
	dist[sid] = 0
	q := &PriorityQueue{[]Vertex{}, make(map[Vertex]int), make(map[Vertex]int64)}
	for _, v := range g.VertexList() {
		if v != sid {
			dist[v] = INFINITY
		}
		prev[v] = UNINITVALUE
		q.AddWithPriority(v, dist[v])
	}
	for len(q.items) != 0 {
		u := heap.Pop(q).(Vertex)
		for _, v := range g.Neighbors(u) {
			alt := dist[u] + g.Weight(u, v)
			if alt < dist[v] {
				dist[v] = alt
				prev[v] = u
				q.Update(v, alt)
			}
		}
	}
	return dist, prev
}

func (g *Graph) Path(v Vertex, prev map[Vertex]Vertex) (pathlist []string) {
	s := g.LabelMap[v]
	pathlist = append(pathlist, s)
	for prev[v] >= 0 {
		v = prev[v]
		pathlist = append([]string{g.LabelMap[v]}, pathlist...)
	}
	return pathlist
}
