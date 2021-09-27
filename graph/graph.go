package graph

// edgeLength specifies the length of edges in the graph.
// It is unclear whether any value other than 2 will
// make sense or be useful in any way ...
const edgeLength = 2

type (
	graph struct {

		// V represents a set of nodes (assumed to be non-empty)
		V Vertices

		// E represents a set of edges (unordered pairs of nodes, may be empty)
		E Edges
		I SyncMapper
	}

	Any interface{}
)

// Order returns the number of nodes (vertices)
func (g *graph) Order() int { return len(*g.V) }

// Size returns the number of edges
func (g *graph) Size() int { return len(g.E.list) }

// Valency (or degree) of a vertex is the number of edges that are
// incident to it, where a loop is counted twice.
func Valency() int {
	// TODO - implement this
	return 0
}

// Degree represents the
func (g *graph) Degree() (max int) {

	for _, i := range g.DegreeMap() {
		if i > max {
			max = i
		}
	}

	return max
}

func (g *graph) DegreeMap() (degreeMap map[Node]int) {
	degreeMap = make(map[Node]int, g.Size())

	for _, n := range g.E.list {
		n0 := n.Node0()
		degreeMap[n0] += 1
	}
	return
}

type vertices []*node
