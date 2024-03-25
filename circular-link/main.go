package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

type CircularLink struct {
	Head *Node
	Tail *Node
}

func NewCircularLink() *CircularLink {
	return &CircularLink{}
}

func (c *CircularLink) InsertAtEnd(value int) {
	node := &Node{Value: value}

	if c.Head == nil {
		c.Head = node
		c.Tail = node
		node.Next = c.Head
	} else {
		node.Next = c.Head
		c.Tail.Next = node
		c.Tail = node
	}
}

func (c *CircularLink) Traverse() {
	if c.Head == nil {
		return
	}
	current := c.Head
	for {
		fmt.Println(current.Value)
		current = current.Next
		if current == c.Head {
			break
		}
	}
}
