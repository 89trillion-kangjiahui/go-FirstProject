package stack

import "errors"

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
func (stack Stack) Top() (string, error) {
	if len(stack) == 0 {
		return "", errors.New("Out of index, len is 0")
	}
	return stack[len(stack)-1], nil
}

func (stack *Stack) Pop() (string, error) {
	theStack := *stack
	if len(theStack) == 0 {
		return "", errors.New("Out of index, len is 0")
	}
	value := theStack[len(theStack)-1]
	*stack = theStack[:len(theStack)-1]
	return value, nil
}

func (stack Stack) IsEmpty() bool {
	return len(stack) == 0
}
