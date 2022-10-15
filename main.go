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
	var head *linked.Node;
	// fmt.Printf("head address: %v\n", &head);
	
	linked.InsertNodeAtHead(linked.CreateNode(3), &head);
	linked.InsertNodeAtHead(linked.CreateNode(6), &head);
	linked.InsertNodeAtHead(linked.CreateNode(9), &head);
	
	linked.InsertAtPosition(linked.CreateNode(10), 1, head);
	linked.PrintLinkedList(head);

	linked.FindNode(4, head);

	linked.RemoveNodeAtHead(&head);
	linked.PrintLinkedList(head);


}