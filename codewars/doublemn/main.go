package main

import "fmt"

func combine(arr1, arr2 []int) []int {
	var arr3 = make(map[int]int)
	var ans []int
	for i := 0; i < len(arr1); i++ {
		arr3[arr1[i]] = arr1[i]
	}
	for i := 0; i < len(arr2); i++ {
		arr3[arr2[i]] = arr2[i]
	}
	for _, j := range arr3 {
		ans = append(ans, j)
	}
	return ans
}

func main() {
	var qca, qcb []int
	var n, a int
	fmt.Println("Укажите размер множеств:")
	fmt.Scanln(&n)
	for i := 0; i < n*2; i++ {
		fmt.Scan(&a)
		if i < n {
			qca = append(qca, a)
		} else {
			qcb = append(qcb, a)
		}
	}
	fmt.Println("Ответ:")
	fmt.Println(combine(qca, qcb))
}
