package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Deque struct {
	items []int
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) AddRear(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) AddFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) RemoveFront() int {
	if len(d.items) == 0 {
		return 0 // или можно выбросить ошибку
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *Deque) RemoveRear() int {
	if len(d.items) == 0 {
		return 0 // или можно выбросить ошибку
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item
}

func (d *Deque) Size() int {
	return len(d.items)
}

func main() {
	str := "1 -2 3 -4"
	s := strings.Split(str, " ")
	q := &Deque{}
	numbers := []int{}

	for _, nums := range s {
		num, _ := strconv.Atoi(nums)
		numbers = append(numbers, num)
	}

	for _, n := range numbers {
		if n < 0 {
			q.AddFront(n) // [-1], [-1 2], [-3 -1 2], [-3 -1 2 4], положительные по порядку
		} else {
			q.AddRear(n)
		}
	}

	for !q.IsEmpty() { // отрицательные в конец
		n := q.RemoveRear()
		if n < 0 { // [-1 2 4 -3], [2 4 -3 -1]
			q.AddRear(n)
		} else {
			q.AddFront(n)
			break
		}
	}

	for !q.IsEmpty() { // берем с конца
		n := q.RemoveFront()
		if n < 0 { // -1 -3 2 4
			fmt.Println(n)
		} else {
			q.AddRear(n)
			break
		}
	}

	for !q.IsEmpty() {
		fmt.Println(q.RemoveRear())
	}
}
