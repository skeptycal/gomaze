package graph

type Edges interface {
	Size() int
	First() Edge
	List() []Edge
	Last() Edge
	CanLoop() bool
}

// edges represents a collection of edges
// (unordered pairs of nodes)
// This basic type has only undirected edges and cannot loop
type edges struct {
	// list is a list of pointers to all edges in the graph
	list []*edge
}

type loopEdges struct{ edges }

func (e loopEdges) CanLoop() bool { return true }

type directedEdges struct{ edges }

func (e directedEdges) CanDirect() bool { return true }

func (e edges) Size() int           { return len(e.list) }
func (e edges) First() Edge         { return e.list[0] }
func (e edges) Last() Edge          { return e.list[len(e.list)-1] }
func (e edges) Get(i int) Edge      { return e.list[i] }
func (e edges) Direction() directed { return Undirected }

func (e edges) CanLoop() bool   { return false }
func (e edges) CanDirect() bool { return false }

type (
	Edge interface {
		Node0() Node
		Node1() Node
		Direction() directed
	}
)

// edge represents an edge between nodes. It can by either directed or undirected.
type edge struct {
	id int

	// nodes is an array of 2 nodes. Element zero is
	// considered the 'first' node and element one is
	// considered the 'second' node.
	nodes *[edgeLength]node
}

func (e *edge) Node0() Node         { return &e.nodes[0] }
func (e *edge) Node1() Node         { return &e.nodes[edgeLength-1] }
func (e *edge) Direction() directed { return Undirected }

type loop struct{ edge }

func (e *loop) Node0() Node         { return &e.nodes[0] }
func (e *loop) Node1() Node         { return e.Node0() }
func (e *loop) Direction() directed { return Undirected }

type directedEdge struct {
	edge

	// directed specifies whether the edge is directed or not.
	//
	// -  'undirected' (0) means it is undirected.
	//
	// - 'forward' (1) means it is directed from node 0 to node 1.
	//
	// - 'backward' (-1) means it is directed from node 1 to node 0.
	directed directed
}

func (e *directedEdge) Direction() directed { return e.directed }

type directed int

const (
	Backward directed = iota - 1
	Undirected
	Forward
)
