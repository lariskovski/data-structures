package stack

import (
	"math"
)

// Max number of elements stack will hold
const stack_length int = 5;

// stack_empty represents the minium value for a 64 bit int
// it'll only be returned by pop if the stack is empty
const stack_empty = math.MinInt64;

// top will keep track of number of elements on the stack
var top int = -1;

var head *Node = nil;

type Node struct {
	value int;
	next *Node;
}

func createNode(value int) *Node{
	var node *Node = new(Node);
	node.value = value;
	node.next = nil;
	return node;
}

func insertNodeAtHead(node *Node, head **Node){
	node.next = *head;
	*head = node;
}

func removeNodeAtHead(head **Node) *Node{
	var node *Node = *head;
	*head = (*head).next;
	return node;
}

func Push(value int) bool{
	if (top < stack_length - 1) {
		top ++;
		insertNodeAtHead(createNode(value), &head)
		return true;
	}
	return false;
}

func Pop() int{
	if (top < 0) {
		top --;
		return removeNodeAtHead(&head).value;
	}
	return stack_empty;
}
