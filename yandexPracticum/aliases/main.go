package main

import "fmt"

func main() {
	var a = "string"
	var b = "my string"
	a = b
	fmt.Println(a)
}

type Mystring = string
