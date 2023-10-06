package main

import (
	"fmt"
	"sort"
)

func combine(n int) []int {
	var arr3 = make(map[int]int)
	var ans, arr1, arr2 []int
	var a int
	for i := 0; i < n*2; i++ {
		fmt.Scan(&a)
		if i < n {
			arr1 = append(arr1, a)
		} else {
			arr2 = append(arr2, a)
		}
	}
	fmt.Println("Ответ:")
	for i := 0; i < len(arr1); i++ {
		arr3[arr1[i]] = arr1[i]
	}
	for i := 0; i < len(arr2); i++ {
		arr3[arr2[i]] = arr2[i]
	}
	for _, j := range arr3 {
		ans = append(ans, j)
	}
	sort.Ints(ans)
	return ans
}

func main() {
	var n int
	fmt.Println("Укажите размер множеств:")
	fmt.Scanln(&n)
	fmt.Println(combine(n))
}
