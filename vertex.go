package gograph

//AddVertex is used to add vertex in graph
//Users also could directly modify IDs/Lables map
//but don't forget to modify VertexNum
func (g *Graph) AddVertex(label string) {
	g.VertexNum++
	id := Vertex(g.VertexNum)
	g.VertexMap[label] = id
	g.LabelMap[id] = label
}
