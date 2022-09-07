package main

type Node struct {
	data int
	next *Node
}

type SinglyLinkedList struct {
	head *Node
	tail *Node
}
