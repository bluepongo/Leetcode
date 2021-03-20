package main

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {
	stack := []int{}
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack, val)
		} else {
			num1, num2 := stack[len(stack)-2], stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+":
				stack = append(stack, num1+num2)
			case "-":
				stack = append(stack, num1-num2)
			case "*":
				stack = append(stack, num1*num2)
			case "/":
				stack = append(stack, num1/num2)
			}
		}
	}
	return stack[0]
}

func evalRPN002(tokens []string) int {
	stack := make([]int, (len(tokens)+1)/2)
	index := -1
	for _, token := range tokens {
		val, err := strconv.Atoi(token)
		if err == nil {
			index++
			stack[index] = val
		} else {
			index--
			switch token {
			case "+":
				stack[index] += stack[index+1]
			case "-":
				stack[index] -= stack[index+1]
			case "*":
				stack[index] *= stack[index+1]
			case "/":
				stack[index] /= stack[index+1]
			}
		}
	}
	return stack[index]
}

func main() {
	tokens := []string{"2", "1", "+", "3", "*"}
	res := evalRPN002(tokens)
	fmt.Println(res)
}
