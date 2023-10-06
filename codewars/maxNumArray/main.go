package main

import "fmt"

func maxNumArray(n int) int {
	var t int
	mas := make([]int, n)
	fmt.Println("Введите значения массива через пробел")
	for i := 0; i < n; i++ {
		fmt.Scan(&t)
		mas[i] = t
	}
	max := mas[0]
	for _, j := range mas {
		if j > max {
			max = j
		}
	}
	return max
}

func main() {
	var n int
	fmt.Println("Введите размер массива")
	fmt.Scanln(&n)
	fmt.Print(maxNumArray(n))
}
