package main

import (
	"fmt"
	"data-structures/linked"
	"data-structures/stack"
)

func main(){

	// =====================
	// Stack Example
	// =====================
	stack.Push(5);
	stack.Push(3);
	stack.Push(1);

	fmt.Printf("%v\n", stack.Pop())
	fmt.Printf("%v\n", stack.Pop())
	fmt.Printf("%v\n", stack.Pop())

	// =====================
	// Linked List Example
	// =====================
	linked.Insert(3)
	linked.Insert(6)
	linked.Insert(9)
	linked.Insert(10, 1)
	linked.PrintLinkedList();

	linked.FindNode(4);

	linked.RemoveNodeAtHead();
	linked.PrintLinkedList();

	linked.Insert(1)

}