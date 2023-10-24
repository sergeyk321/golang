package main

import "fmt"

func main() {
	var n, row, place, price int
	fmt.Scanln(&n)
	ans := make(map[[2]int]map[int]bool)
	for i := 0; i < n; i++ {
		fmt.Scanln(&row, &place, &price)
		key := [2]int{row, place}
		if ans[key] == nil {
			ans[key] = make(map[int]bool)
		}
		ans[key][price] = true
	}
	for key, value := range ans {
		fmt.Printf("%d %d - %d\n", key[0], key[1], len(value))
	}
}
