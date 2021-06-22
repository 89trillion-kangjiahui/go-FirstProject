package util

import (
	"strconv"
	"unicode"

	. "project2/stack"
)

func GetResult(postfix []string) int {
	var curStack Stack
	for _, v := range postfix {
		curChar := v
		if i, ero := strconv.Atoi(v); ero == nil {
			curStack.Push(strconv.Itoa(i))
		} else {
			v1, _ := curStack.Pop()
			v2, _ := curStack.Pop()
			num1, _ := strconv.Atoi(v1)
			num2, _ := strconv.Atoi(v2)
			switch curChar {
			case "+":
				curStack.Push(strconv.Itoa(num2 + num1))
			case "-":
				curStack.Push(strconv.Itoa(num2 - num1))
			case "*":
				curStack.Push(strconv.Itoa(num2 * num1))
			case "/":
				curStack.Push(strconv.Itoa(num2 / num1))
			}
		}
	}
	ret, _ := curStack.Top()
	result, _ := strconv.Atoi(ret)
	return result
}

func MixToPost(exp string) []string {
	var curStack Stack
	var prefix []string
	expLen := len(exp)
	for i := 0; i < expLen; i++ {
		char := string(exp[i])
		switch char {
		case " ":
			continue
		case "(":
			curStack.Push("(")
		case ")":
			for !curStack.IsEmpty() {
				top, _ := curStack.Top()
				if top == "(" {
					curStack.Pop()
					break
				}
				prefix = append(prefix, top)
				curStack.Pop()
			}
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			j := i
			digit := ""
			for ; j < expLen && unicode.IsDigit(rune(exp[j])); j++ {
				digit += string(exp[j])
			}
			prefix = append(prefix, digit)
			i = j - 1
		default:
			for !curStack.IsEmpty() {
				top, _ := curStack.Top()
				if top == "(" || isMore(top, char) {
					break
				}
				prefix = append(prefix, top)
				curStack.Pop()
			}
			curStack.Push(char)
		}
	}
	for !curStack.IsEmpty() {
		data, _ := curStack.Pop()
		prefix = append(prefix, data)
	}
	return prefix
}

func isMore(top, char string) bool {
	switch top {
	case "+", "-":
		if char == "*" || char == "/" {
			return true
		}
	case "(":
		return true
	}
	return false
}
