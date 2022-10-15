package linked

import (
	"fmt"
)

type Node struct {
	value int;
	next *Node;
}

var head *Node;

func PrintLinkedList(){
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

func createNode(value int) *Node{
	var node *Node = new(Node);
	node.value = value;
	node.next = nil;
	return node;
}

func insertNodeAtHead(node *Node){
	// fmt.Printf("insert %v at head (currently %v)\n", node.value, (*head).value);
	// fmt.Printf("head address: %v\n", head);
	var temp_head **Node = &head;
	node.next = *temp_head;
	*temp_head = node;
}

func RemoveNodeAtHead(){
	var temp_head **Node = &head;
	*temp_head = (*temp_head).next;
}


func FindNode(value int){
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

func insertAtPosition(node *Node, position uint){
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

func Insert(value int, position ...uint) {
	// Implement error throw if position > 1
	if (position == nil) {
		insertNodeAtHead(createNode(value))
	} else {
		insertAtPosition(createNode(value), position[0]);
	}
}