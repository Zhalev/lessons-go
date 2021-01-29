package hw_3

type Node struct {
	prev *Node
	next *Node
	data interface{}
}
type List struct {
	FirstNode *Node
	LastNode  *Node
}

func (l *List) Len() int {
	nodeTemp := l.FirstNode
	i := 0
	for nodeTemp != nil {
		i++
		nodeTemp = nodeTemp.next
	}
	return i
}
func (l *List) First() *Node {
	return l.FirstNode
}
func (l *List) Last() *Node {
	return l.LastNode
}
func (l *List) PushBack(v interface{}) {
	node := &Node{data: v, prev: nil, next: nil}

	if l.LastNode == nil {
		l.LastNode = node
	} else {
		node.prev = l.LastNode
		l.LastNode.prev = node
		l.LastNode = node

	}
}
func (l *List) PushFront(v interface{}) {

	node := &Node{data: v, prev: nil, next: nil}

	if l.FirstNode == nil {
		l.FirstNode = node
	} else {
		node.next = l.FirstNode
		l.FirstNode.prev = node
		l.FirstNode = node

	}

}
func (l *List) Remove(n *Node) {
	if n.prev == nil {
		l.FirstNode = n.next
	} else {
		n.prev.next = n.next
	}
	if n.next == nil {
		l.LastNode = n.prev
	} else {
		n.next.prev = n.prev
	}
	/*	node := l.FirstNode
		for i != node {
			node = l.LastNode.next
		}
		if node == i {
			node.next.prev = node.prev.next
		}
	*/
}

func (n *Node) Value() interface{} {
	return n.data
}
func (n *Node) Next() *Node {
	return n.next
}
func (n *Node) Prev() *Node {
	return n.prev
}
