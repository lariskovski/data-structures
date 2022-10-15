package stack

const stack_length int = 5;
const empty int = -1;

var stack [stack_length]int;
var top int = empty;

func Push(value int) bool{
	if (top < stack_length - 1) {
		top ++;
		stack[top] = value;
		return true;
	}
	return false;
}

func Pop() int{
	if (top != empty) {
		var tmp int = stack[top]
		top --;
		return tmp;
	}
	return -1;
}
