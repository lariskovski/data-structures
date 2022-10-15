package linked

import (
	"fmt"
)

type Node struct {
	value int;
	next *Node;
}

func PrintLinkedList(head *Node){
	var temp *Node = head;
	var count uint = 0;
	for  {
		fmt.Printf("Node %v: %v\n", count, temp.value);
		if temp.next == nil{
			break
		}
		temp = temp.next;
		count ++;
	}
}

func CreateNode(value int) *Node{
	var node *Node = new(Node);
	node.value = value;
	node.next = nil;
	return node;
}

func InsertNodeAtHead(node *Node, head **Node){
	// fmt.Printf("insert %v at head (currently %v)\n", node.value, (*head).value);
	// fmt.Printf("head address: %v\n", head);
	node.next = *head;
	*head = node;
}

func RemoveNodeAtHead(head **Node){
	*head = (*head).next;
}


func FindNode(value int, head *Node){
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
	fmt.Printf("Int %v not found on linked list.\n", value);
}

func InsertAtPosition(node *Node, position uint,head *Node){
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