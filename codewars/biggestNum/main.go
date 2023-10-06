package main

import (
	"fmt"
	"sort"
)

func NextBigger(n int) int {
	w := fmt.Sprint(n)
	var mas []int
	ans := 0
	for _, v := range w {
		mas = append(mas, int(v)-48)
	}
	sort.Ints(mas)
	for i := len(w) - 1; i >= 0; i-- {
		fmt.Println(mas[i])
		ans = ans*10 + mas[i]
	}
	// your code here
	return ans
}

func main() {
	var n int
	fmt.Scanln(&n)
	fmt.Print(NextBigger(n))
}
