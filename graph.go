package graph_playground

type Link struct {
	Vertice *VerticeLinked
	Next    *Link
}

func (l Link) GetLast() *Link {
	if l.Next == nil {
		return &l
	}
	return l.Next.GetLast()
}

func (l Link) Append(v *VerticeLinked) {
	last := l.GetLast()
	last.Next = &Link{
		Vertice: v,
		Next:    nil,
	}
}

type VerticeLinked struct {
	// Support multiple relationship types via a map of linked lists
	Edges map[string]*Link
	Name  string
}

func (v VerticeLinked) AddEdge(v_new *VerticeLinked, relationship string) {
	if v.Edges[relationship] == nil {
		v.Edges[relationship] = &Link{
			Vertice: v_new,
			Next:    nil,
		}
	} else {
		v.Edges[relationship].Append(v_new)
	}

}

type Graph struct {
	Ceiling  uint
	Vertices map[uint]VerticeLinked
}

func InitializeGraph() Graph {
	vertices := make(map[uint]VerticeLinked)
	return Graph{
		Ceiling:  0,
		Vertices: vertices,
	}
}

func (g *Graph) AddVertice(v VerticeLinked) uint {
	g.Ceiling = g.Ceiling + 1
	g.Vertices[g.Ceiling] = v
	return g.Ceiling
}

func (g Graph) AddEdge(from uint, to uint, fromRelationship string, toRelationship string) {
	// Find from
	f := g.Vertices[from]
	t := g.Vertices[to]
	f.AddEdge(&t, fromRelationship)
	t.AddEdge(&f, toRelationship)
}

func InitializeVertice(name string) VerticeLinked {
	return VerticeLinked{
		Edges: make(map[string]*Link),
		Name:  name,
	}
}
