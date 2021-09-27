package graph

// Node represents a node in the graph. For basic nodes,
// only Get and Set are supported.
//
// More specialized implementations of node support
// Name, Id, Next, and/or Prev
type Node interface {
	// Get returns the data value of the node.
	Get() Any

	// Set sets the data value of the node.
	Set(value Any)

	// Name returns the name of the node, if any.
	Name() string

	// Id returns the id of the node, if any.
	Id() Any

	// Next returns the next node if the the set of
	// nodes forms a linked list, otherwise nil.
	Next() Node

	// Prev returns the previous node if the the set of
	// nodes forms a double linked list, otherwise nil.
	Prev() Node
}

type nodePair [edgeLength]*node

type node struct {
	data Any // data is the payload of the node.
}

func (n *node) Get() Any      { return n.data }
func (n *node) Set(value Any) { n.data = value }

func (n *node) Name() string      { return "" }
func (n *node) Id() Any           { return 0 }
func (n *node) Next() Node        { return nil }
func (n *node) Prev() Node        { return nil }
func (n *node) getRawNode() *node { return n }
