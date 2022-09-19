package main

import "fmt"

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head   *Node
	tail   *Node
	length int
}

/* Constructors */
// Singly Linked List Constructor
func NewSinglyLinkedList() *SinglyLinkedList {
	return &SinglyLinkedList{}
}

// Node constructor
func NewNode(data int) *Node {
	return &Node{
		data: data,
	}
}

/*Methods*/

// Traverse the list
func (sll *SinglyLinkedList) Traversal() {
	pointer := sll.head
	for pointer != nil {
		fmt.Println(pointer.data)
		pointer = pointer.next
	}
}

/* Add nodes to the list */
// Add a node at the end of the list
func (sll *SinglyLinkedList) AddNodeEnd(data int) {
	node := NewNode(data)
	if sll.head == nil {
		sll.head = node
		sll.tail = node
	} else {
		sll.tail.next = node
		sll.tail = node
	}
}

/* Add a node at the top of the list */
func (sll *SinglyLinkedList) AddNodeTop(data int) {
	node := NewNode(data)
	if sll.head == nil {
		sll.head = node
		sll.tail = node
	} else {
		node.next = sll.head
		sll.head = node
	}
	sll.length++
}

/* Add a node at a specific position */
func (sll *SinglyLinkedList) AddNodeSpecify(data int, position int) {
	node := NewNode(data)
	if position == 0 {
		sll.AddNodeTop(node.data)
	} else {
		pointer := sll.head
		for i := 0; i < position-1; i++ {
			pointer = pointer.next
		}
		after := pointer.next
		node.next = after
		pointer.next = node
	}
	sll.length++
}

/* Remove nodes from the list */
// Remove head from list
func (sll *SinglyLinkedList) RemoveHead() {
	if sll.head == nil || sll.tail == nil {
		fmt.Println("Nothing to delete")
	} else {
		if sll.head == sll.tail {
			sll.head = sll.tail.next
			sll.tail = nil
		} else {
			temp := sll.head
			sll.head = temp.next
			*temp = Node{}
		}
		sll.length--
	}
}

// Remove tail from the list
func (sll *SinglyLinkedList) RemoveTail() {
	if sll.head == nil || sll.tail == nil {
		fmt.Println("Nothing to delete")
	} else {
		if sll.head == sll.tail {
			sll.head = sll.head.next
			sll.tail = nil
		} else {
			pointer := sll.head
			for pointer.next.next != nil {
				pointer = pointer.next
			}
			sll.tail = &Node{}
			pointer.next = nil
			sll.tail = pointer
		}
		sll.length--
	}
}

func (sll *SinglyLinkedList) RemoveSpecify(position int) {

}
