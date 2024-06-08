package main

import "fmt"

type Node[T any] struct {
	val      T
	nextNode *Node[T]
}

var head *Node[any]
var tail *Node[any]

func (n *Node[T]) AddElement(data T) {

	node := &Node[any]{val: data}
	if head == nil {
		head = node
		tail = node
		return
	}

	tail.nextNode = node
	tail = node

}

func (n *Node[T]) PrintNode() {
	temp := head

	for temp != nil {
		fmt.Println(temp.val)
		temp = temp.nextNode
	}
}

func main() {
	n := Node[int]{}
	n.AddElement(12)
	n.AddElement(452)
	n.AddElement(45)
	n.AddElement(90)
	n.PrintNode()
}
