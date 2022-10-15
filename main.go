package main

import (
	"fmt"
)

type Node struct {
	value int;
	next *Node;
}

func printLinkedList(head *Node){
	var temp *Node = head;
	var count uint = 0;
	for {
		fmt.Printf("Node %v: %v\n", count, temp.value);
		if temp.next == nil{
			break
		}
		temp = temp.next;
		count ++;
	}
}

func createNode(value int) *Node{
	var node *Node = new(Node);
	node.value = value;
	node.next = nil;
	return node;
}

func insertNodeAtHead(node *Node, head **Node){
	// fmt.Printf("insert %v at head (currently %v)\n", node.value, (*head).value);
	// fmt.Printf("head address: %v\n", head);
	node.next = *head;
	*head = node;
}

func removeNodeAtHead(head **Node){
	*head = (*head).next;
}


func findNode(value int, head *Node){
	var temp *Node = head;
	var count uint = 0;
	for (temp.next != nil) {
		if temp.value == value{
			fmt.Printf("Found %v at position %v\n", value, count);
			return;
		} else {
			temp = temp.next;
			count ++;
		}
	}
	fmt.Printf("Int %v not found on linked list.", value);
}


func insertAtPosition(node *Node, position uint,head *Node){
	var temp *Node = head;
	var count uint = 0;
	for (temp.next != nil) {
		if (count == position - 1) {
			node.next = temp.next;
			temp.next = node;
			return;
		} else {
			temp = temp.next;
			count ++;
		}
	}
	fmt.Printf("Position %v out of range. Linked List size is %v", position, count);
}

func main(){

	var head *Node;
	// fmt.Printf("head address: %v\n", &head);
	
	insertNodeAtHead(createNode(24), &head);
	insertNodeAtHead(createNode(99), &head);
	insertNodeAtHead(createNode(9), &head);
	insertNodeAtHead(createNode(3), &head);

	insertAtPosition(createNode(10), 1, head);

	printLinkedList(head);

	// findNode(4, head);
	
	// removeNodeAtHead(&head);
	// printLinkedList(head);


}