package hierarchy

/*
 Hierarchy pattern library helper
*/

type Node[T comparable] struct {
	// Value for current node
	Value T

	// Childrens for the node
	Children []*Node[T]
}

func NewHierarchy[T comparable](value T, children ...*Node[T]) *Node[T] {
	return &Node[T]{
		Value:    value,
		Children: children,
	}
}

// AddChild adds a child to the current node
func (n *Node[T]) AddChild(child *Node[T]) *Node[T] {
	// Returns newly added child
	n.Children = append(n.Children, child)
	return child
}

// FindChild finds the first child of the node matching the value
func (n *Node[T]) FindChild(value T) *Node[T] {
	for _, child := range n.Children {
		if child.Value == value {
			return child
		}
	}
	return nil
}
