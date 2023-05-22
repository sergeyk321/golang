package main

import (
	"errors"
	"fmt"
	"strings"
)

func createStack() []string {
	stack := make([]string, 0)
	return stack
}

func checkEmpty(stack []string) bool {
	return len(stack) == 0
}

func push(stack *[]string, item string) {
	*stack = append(*stack, item)
}

func pop(stack *[]string) (string, error) {
	if checkEmpty(*stack) {
		return "", errors.New("Стек пуст")
	}
	item := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return item, nil
}

func main() {
	s := "Преступление и наказание, Три товарища, Униженные и оскорбленные, Портрет Дориана Грея, 20000 лье под водой"
	str := strings.Split(s, ", ")
	st := createStack()
	for _, c := range str {
		push(&st, c)
	}
	for !checkEmpty(st) {
		item, _ := pop(&st)
		fmt.Printf("%s\n", item)
	}
}
