package main

import "fmt"

type Stack struct {
	items []int
}

func (s *Stack) Push(item int) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {
	if len(s.items) == 0 {
		return -1 // или можно выбросить ошибку
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return item
}

func main() {
	a := &Stack{}
	b := &Stack{}
	c := &Stack{}

	n := 5
	for i := n; i > 0; i-- {
		a.Push(i)
	}

	if n%2 == 0 {
		for len(c.items) != n {
			circular_queue(a, b)
			circular_queue(a, c)
			circular_queue(b, c)
			fmt.Println("Башни:")
			fmt.Println(*a)
			fmt.Println(*b)
			fmt.Println(*c)
		}
	} else {
		for len(c.items) != n {
			circular_queue(a, b)
			for len(a.items) == 0 && len(b.items) == 0 {
				break
			}
			circular_queue(a, c)
			circular_queue(c, b)
			fmt.Println("Башни:")
			fmt.Println(*a)
			fmt.Println(*b)
			fmt.Println(*c)
		}
	}
	fmt.Println("Вывод:")
	fmt.Println(*a)
	fmt.Println(*b)
	fmt.Println(*c)
}

func circular_queue(a, b *Stack) {
	if len(a.items) == 0 && len(b.items) == 0 {
		return
	}
	if len(a.items) > 0 && len(b.items) == 0 {
		b.Push(a.Pop())
	} else if len(a.items) == 0 && len(b.items) > 0 {
		a.Push(b.Pop())
	} else if len(a.items) > 0 && len(b.items) > 0 {
		if a.items[len(a.items)-1] < b.items[len(b.items)-1] {
			b.Push(a.Pop())
		} else {
			a.Push(b.Pop())
		}
	}
}
