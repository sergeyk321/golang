package main

import (
	"fmt"
)

func createStack() []rune {
	stack := []rune{}
	return stack
}

func checkEmpty(stack []rune) bool {
	return len(stack) == 0
}

func push(stack *[]rune, item rune) {
	*stack = append(*stack, item)
}

func pop(stack *[]rune) interface{} {
	if checkEmpty(*stack) {
		return "Стек пуст"
	}
	item := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return item
}

func main() {
	s := "asd23 ./!/3das"
	answer := ""

	numbers := createStack()
	letters := createStack()
	symbols := createStack()

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			push(&numbers, ch)
		} else if (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') {
			push(&letters, ch)
		} else {
			push(&symbols, ch)
		}
	}
	for !checkEmpty(numbers) {
		answer += string(pop(&numbers).(rune)) //интерфейс в руну
	}
	for !checkEmpty(letters) {
		answer += string(pop(&letters).(rune))
	}
	for !checkEmpty(symbols) {
		answer += string(pop(&symbols).(rune))
	}

	fmt.Println(answer)
}
