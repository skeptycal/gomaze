package graph

func newDoubleLinkedListNode(next, prev Node) Node {
	return &doubleLinkedListNode{linkedListNode{node{}, next}, prev}
}

type doubleLinkedListNode struct {
	linkedListNode
	prev Node
}

func (n *doubleLinkedListNode) Prev() Node        { return n.prev }
func (n *doubleLinkedListNode) SetPrev(prev Node) { n.prev = prev }
