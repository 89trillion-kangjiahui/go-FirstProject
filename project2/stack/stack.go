package stack

type Stack []string

//统计栈中元素的数量
func (stack Stack) Len() int {
	return len(stack)
}

//统计栈的容量
func (stack Stack) Cap() int {
	return cap(stack)
}

//将元素押入栈
func (stack *Stack) Push(value string) {
	*stack = append(*stack, value)
}

//获取栈顶元素
func (stack Stack) Top() string {
	return stack[len(stack)-1]
}

func (stack *Stack) Pop() string {
	theStack := *stack
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
