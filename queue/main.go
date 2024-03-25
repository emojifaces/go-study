package main

import "fmt"

type Queue struct {
	elements []int
}

func NewQueue() *Queue {
	return &Queue{
		elements: []int{},
	}
}

// Enqueue 向队列中添加一个元素
func (q *Queue) Enqueue(element int) {
	q.elements = append(q.elements, element)
}

// Dequeue 从队列中移除并返回队首元素
// 如果队列为空，则返回-1（或其他错误值/消息）
func (q *Queue) Dequeue() int {
	if len(q.elements) == 0 {
		fmt.Println("Queue is empty")
		return -1 // 或者可以抛出一个错误
	}
	front := q.elements[0]
	q.elements = q.elements[1:]
	return front
}

// Peek 返回队首元素但不移除
// 如果队列为空，则返回-1（或其他错误值/消息）
func (q *Queue) Peek() int {
	if len(q.elements) == 0 {
		fmt.Println("Queue is empty")
		return -1 // 或者可以抛出一个错误
	}
	return q.elements[0]
}

// IsEmpty 检查队列是否为空
func (q *Queue) IsEmpty() bool {
	return len(q.elements) == 0
}
