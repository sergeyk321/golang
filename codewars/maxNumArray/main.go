package main

import "fmt"

func maxNumArray(arr []int) int {
	max := arr[0]
	for _, j := range arr {
		if j > max {
			max = j
		}
	}
	return max
}

func main() {
	fmt.Print(maxNumArray([]int{1, 5, 7, 10, 6}))
}
