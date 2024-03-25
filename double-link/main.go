package main

import "fmt"

type Node struct {
	Value int
	Prev  *Node
	Next  *Node
}

type DoubleLink struct {
	Head *Node
	Tail *Node
}

func NewDoubleLink() *DoubleLink {
	return &DoubleLink{}
}

func (d *DoubleLink) InsertAtEnd(value int) {
	node := &Node{Value: value}
	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		d.Tail.Next = node
		node.Prev = d.Tail
		d.Tail = node
	}
}

func (d *DoubleLink) InsertAtFront(value int) {
	node := &Node{Value: value}
	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.Next = d.Head
		d.Head.Prev = node
		d.Head = node
	}
}

func (d *DoubleLink) TraverseForward() {
	current := d.Head
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Next
	}
	fmt.Println()
}

func (d *DoubleLink) TraverseBackward() {
	current := d.Tail
	for current != nil {
		fmt.Print(current.Value, " ")
		current = current.Prev
	}
	fmt.Println()
}
