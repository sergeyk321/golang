package main

import (
	"fmt"
	"strconv"
	//"unicode"
)

func SumMix(arr []any) int {
	var sum int
	for _, j := range arr {
		r := fmt.Sprint(j)
		t, _ := strconv.Atoi(r)
		sum += t
	}
	return sum
}

func main() {
	fmt.Print(SumMix([]any{"2", 3, 4, "5", 4}))
}
