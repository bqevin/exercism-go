package linkedlist

import (
    "errors"
)

// Define List and Node types here.
type Node struct {
	Left  *Node
	Value any
	Right *Node
}
type List struct {
	Head  *Node
	Tail  *Node
	Items []*Node
}

// Note: The tests expect Node type to include an exported field with name Value to pass.

func NewList(elements ...any) *List {
	if len(elements) == 1 {
		if s, ok := elements[0].([]any); ok {
			elements = s
		}
	}
	if len(elements) == 0 {
		return &List{Items: []*Node{}}
	}
	items := make([]*Node, len(elements))
	for i, v := range elements {
		items[i] = &Node{Value: v}
		if i > 0 {
			items[i-1].Right = items[i]
			items[i].Left = items[i-1]
		}
	}
	return &List{
		Head:  items[0],
		Tail:  items[len(items)-1],
		Items: items,
	}
}

func (n *Node) Next() *Node {
	return n.Right
}

func (n *Node) Prev() *Node {
	return n.Left
}

// Unshift adds to front.
func (l *List) Unshift(v any) {
	n := &Node{Value: v}
	if l.Head == nil {
		l.Head, l.Tail = n, n
		l.Items = []*Node{n}
		return
	}
	n.Right = l.Head
	l.Head.Left = n
	l.Head = n
	// Update slice – prepend
	l.Items = append([]*Node{n}, l.Items...)
}

// Push adds to end.
func (l *List) Push(v any) {
	n := &Node{Value: v}
	if l.Tail == nil {
		l.Head, l.Tail = n, n
		l.Items = []*Node{n}
		return
	}
	l.Tail.Right = n
	n.Left = l.Tail
	l.Tail = n
	l.Items = append(l.Items, n)
}

// Shift removes first.
func (l *List) Shift() (any, error) {
	if l.Head == nil {
		return 0, errors.New("empty")
	}
	val := l.Head.Value
	if l.Head == l.Tail {
		l.Reset()
		return val, nil
	}
	newHead := l.Head.Right
	newHead.Left = nil
	l.Head = newHead
	l.Items = l.Items[1:] // slice update
	return val, nil
}

// Pop removes last.
func (l *List) Pop() (any, error) {
	if l.Tail == nil {
		return 0, errors.New("empty")
	}
	val := l.Tail.Value
	if l.Head == l.Tail {
		l.Reset()
		return val, nil
	}
	newTail := l.Tail.Left
	newTail.Right = nil
	l.Tail = newTail
	l.Items = l.Items[:len(l.Items)-1]
	return val, nil
}

// Reverse reverses the list.
func (l *List) Reverse() {
	n := len(l.Items)
	if n <= 1 {
		return
	}
	// reverse slice
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		l.Items[i], l.Items[j] = l.Items[j], l.Items[i]
	}
	// relink nodes based on new slice order
	for i := 0; i < n; i++ {
		if i > 0 {
			l.Items[i].Left = l.Items[i-1]
		} else {
			l.Items[i].Left = nil
		}
		if i < n-1 {
			l.Items[i].Right = l.Items[i+1]
		} else {
			l.Items[i].Right = nil
		}
	}
	l.Head = l.Items[0]
	l.Tail = l.Items[n-1]
}

func (l *List) First() *Node {
	return l.Head
}

func (l *List) Last() *Node {
	return l.Tail
}

func (l *List) Count() int {
	return len(l.Items)
}

func (l *List) Reset() {
	l.Head, l.Tail = nil, nil
	l.Items = []*Node{}
}

// Delete removes the first node in a list with a given value.
// Returns true if a node was removed.
// Delete removes first occurrence of v.
func (l *List) Delete(v any) bool {
	for i, node := range l.Items {
		if node.Value == v {
			// remove from links
			if node.Left != nil {
				node.Left.Right = node.Right
			} else {
				l.Head = node.Right
			}
			if node.Right != nil {
				node.Right.Left = node.Left
			} else {
				l.Tail = node.Left
			}
			// remove from slice
			l.Items = append(l.Items[:i], l.Items[i+1:]...)
			// isolate node
			node.Left, node.Right = nil, nil
			return true
		}
	}
	return false
}
