package graph

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

type idNode struct {
	node
	// id is the unique identifier of the node.
	// (could just use a pointer instead ...)
	id int
}

func (n *idNode) Id(id int) Any { return n.id }

// func (n *idNode) getRawNode() *idNode { return n }

type namedNode struct {
	node

	// name is needed for some implementations
	name string
}

func (n *namedNode) Name() string { return n.name }

// func (n *namedNode) getRawNode() *namedNode { return n }
