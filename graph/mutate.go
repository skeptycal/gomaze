package graph

// Mutate changes the node into a new node type
func Mutate(n node, kind NodeType) Node {

	switch kind {
	case BasicNode:
		return &n
	case LinkedListNode:
		return &linkedListNode{n, nil}
	case DoubleLinkedListNode:
		return &doubleLinkedListNode{linkedListNode{n, nil}, nil}
	default:
		return nil
	}
}
