package main

import "fmt"

func main() {
	var n int
	fmt.Scanln(&n)
	var s1, s2 int
	var line, mas []int
	ans := make([][]int, n)
	for i := 0; i < n; i++ {
		fmt.Scanln(&s1, &s2)
		line = append(line, s1)
		line = append(line, s2)
		ans[i] = line
		line = nil
	}
	for _, j := range ans {
		mas = append(mas, j[0]+j[1])
	}
	for _, j := range mas {
		fmt.Println(j)
	}
}
