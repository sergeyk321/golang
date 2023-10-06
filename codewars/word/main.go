package main

import "fmt"

func alphanumeric(str string) bool {
	for _, i := range str {
		fmt.Println(i)
	}
	return true
}

func main() {
	fmt.Println(alphanumeric("abc"))
}
