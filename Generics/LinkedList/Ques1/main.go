package main

import "fmt"

type Node[T any] struct {
	data     T
	nextNode *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
	tail *Node[T]
}

func (l *LinkedList[T]) AddNode(val T) {
	node := &Node[T]{data: val}
	if l.head == nil {
		l.head = node
		l.tail = node
		return

	}
	l.tail.nextNode = node
	l.tail = node
}

func (l *LinkedList[T]) PrintList() {
	temp := l.head
	for temp != nil {
		fmt.Print(temp.data, " -> ")
		temp = temp.nextNode
	}
	fmt.Println()
}

// remove fromm start
func (l *LinkedList[T]) RemoveElement() T {
	if l.head == nil {
		var z T
		return z
	}
	remove := l.head.data
	l.head = l.head.nextNode
	return remove
}

func main() {
	l := LinkedList[int]{}
	for i := 0; i < 5; i++ {
		l.AddNode(i + 10)
	}

	l.PrintList()
	for i := 0; i < 5; i++ {
		fmt.Println(l.RemoveElement())
	}

	l.PrintList()
}
