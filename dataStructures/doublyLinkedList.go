package dataStructures

type DLLNode struct {
	val  int
	prev *DLLNode
	next *DLLNode
}

type DoublyLinkedList struct {
	head *DLLNode
	tail *DLLNode
}

func NewDLLNode(value int, previous *DLLNode, next *DLLNode) *DLLNode {
	newNode := &DLLNode{
		val:  value,
		prev: previous,
		next: next,
	}
	if previous != nil {
		previous.next = newNode
	}
	if next != nil {
		next.prev = newNode
	}
	return newNode
}

func NewDLL(head *DLLNode, tail *DLLNode) *DoublyLinkedList {
	return &DoublyLinkedList{
		head: head,
		tail: tail,
	}
}

func (dLL *DoublyLinkedList) AddHead(value int) *DoublyLinkedList {
	NewDLLNode(value, nil, dLL.head)
	return dLL
}

func (dLL *DoublyLinkedList) AddTail(value int) *DoublyLinkedList {
	NewDLLNode(value, dLL.tail, nil)
	return dLL
}
