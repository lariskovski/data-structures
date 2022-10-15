package main

import (
	"data-structures/linked"
)

func main(){

	// =====================
	// Linked List Examples
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