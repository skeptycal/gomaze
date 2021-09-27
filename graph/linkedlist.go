package graph

type NodeType int

const (
	BasicNode NodeType = iota
	LinkedListNode
	DoubleLinkedListNode
)

func newLinkedListNode(next Node) Node {
	return &linkedListNode{node{}, next}
}

type linkedListNode struct {
	node
	next Node
}

func (n *linkedListNode) Next() Node        { return n.next }
func (n *linkedListNode) SetNext(next Node) { n.next = next }
