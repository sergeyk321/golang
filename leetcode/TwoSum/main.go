package main

import (
	"fmt"
	"sort"
)

func main() {
	var n, k, target int
	var nums []int
	fmt.Scanln(&n)
	fmt.Scanln(&target)
	fmt.Println("Вводите элементы:")
	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &k)
		nums = append(nums, k)
	}
	fmt.Println(x2(nums, target))
	fmt.Println(x(nums, target))
	fmt.Println(twoPoints(nums, target))
}

// Time: O(n^2) Space: O(1)
func x2(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}

// Time: O(2*n)=O(n) Space: O(n)
func x(nums []int, target int) []int {
	s := make(map[int]int, 0)
	for i, num := range nums {
		s[num] = i //числа из массива : их id
	}
	for i, num := range nums {
		if j, ok := s[target-num]; ok && i != j {
			return []int{i, j}
		} //ищем значения target-num, ведь ключи в s - значения из nums
	}
	return []int{}
}

// Time: O(n), Space: O(1), nums должны быть отсортированы
func twoPoints(nums []int, target int) []int {
	if len(nums) <= 1 {
		return []int{}
	}
	sort.Ints(nums)
	var ptr1, ptr2 = 0, len(nums) - 1
	for ptr1 < ptr2 {
		current := nums[ptr1] + nums[ptr2]
		if current == target {
			return []int{ptr1, ptr2}
		} else if current > target { //если получившаяся сумма больше искомого числа, сдвигаем бОльший укатель влево
			ptr2--
		} else {
			ptr1++
		}
	}
	return []int{}
}
