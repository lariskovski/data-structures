package stack

const stack_length int = 5;
const empty int = -1;
var top int = empty;

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
	// fmt.Printf("insert %v at head (currently %v)\n", node.value, (*head).value);
	// fmt.Printf("head address: %v\n", head);
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
	if (top != empty) {
		top --;
		return removeNodeAtHead(&head).value;
	}
	return -1;
}
