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
	
	// fmt.Printf("head address: %v\n");
	
	linked.InsertNodeAtHead(linked.CreateNode(3));
	linked.InsertNodeAtHead(linked.CreateNode(6));
	linked.InsertNodeAtHead(linked.CreateNode(9));
	
	linked.InsertAtPosition(linked.CreateNode(10), 1);
	linked.PrintLinkedList();

	linked.FindNode(4);

	linked.RemoveNodeAtHead();
	linked.PrintLinkedList();


}