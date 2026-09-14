package dataStructures

type Node struct {
	val  int
	next *Node
}

type LinkedList struct {
	head *Node
}

func newNode(value int, prev *Node) *Node {
	newNode := &Node{
		val:  value,
		next: nil,
	}
	if prev != nil {
		prev.next = newNode
	}
	return newNode
}

func NewLinkedList(head *Node) *LinkedList {
	return &LinkedList{
		head: head,
	}
}

func (l *LinkedList) AddNode(value int) *LinkedList {
	tail := l.head
	for tail != nil && tail.next != nil {
		tail = tail.next
	}
	newNode(value, tail)
	return l
}
