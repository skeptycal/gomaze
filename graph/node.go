package graph

type NodeType int

const (
	BasicNode NodeType = iota
	LinkedListNode
	DoubleLinkedListNode
)

func newNode(data Any, kind NodeType) Node {

	n := &node{data: data}

	switch kind {
	case BasicNode:
		return n
	case LinkedListNode:
		return Mutate(n, kind)
	case DoubleLinkedListNode:
		return &doubleLinkedListNode{linkedListNode{*n, nil}, nil}
	}
	return n
}

func newIdNode(id int) Node {
	return &idNode{node{}, id}
}

func newNamedNode(name string) Node {
	return &namedNode{node{}, name}
}

func newLinkedListNode(next Node) Node {
	return &linkedListNode{node{}, next}
}

func newDoubleLinkedListNode(next, prev Node) Node {
	return &doubleLinkedListNode{linkedListNode{node{}, next}, prev}
}

// Mutate changes the node into a new node type
func Mutate(n Node, kind NodeType) Node {

	switch kind {
	case BasicNode:
		return n
	case LinkedListNode:
		return &linkedListNode{n, nil}
	case DoubleLinkedListNode:
		return newDoubleLinkedListNode(nil, nil)
	}
}

type Node interface {
	Get() Any
	Set(value Any)
	Name() string
	Next() Node
	Prev() Node
}

type nodePair [edgeLength]*node

type node struct {
	data Any // data is the payload of the node.
}

func (n *node) Get() Any      { return n.data }
func (n *node) Set(value Any) { n.data = value }
func (n *node) Name() string  { return "" }
func (n *node) Id() Any       { return 0 }
func (n *node) Next() Node    { return nil }
func (n *node) Prev() Node    { return nil }

type idNode struct {
	node
	// id is the unique identifier of the node.
	// (could just use a pointer instead ...)
	id int
}

func (n *idNode) Id(id int) Any { return n.id }

type namedNode struct {
	node

	// name is needed for some implementations
	name string
}

func (n *namedNode) Name() string { return n.name }

type linkedListNode struct {
	node
	next Node
}

func (n *linkedListNode) Next() Node        { return n.next }
func (n *linkedListNode) SetNext(next Node) { n.next = next }

type doubleLinkedListNode struct {
	linkedListNode
	prev Node
}

func (n *doubleLinkedListNode) Prev() Node        { return n.prev }
func (n *doubleLinkedListNode) SetPrev(prev Node) { n.prev = prev }
