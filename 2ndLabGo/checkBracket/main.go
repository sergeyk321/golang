package main

import "fmt"

type Deque struct {
	items []string
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) AddRear(item string) {
	d.items = append(d.items, item)
}

func (d *Deque) AddFront(item string) {
	d.items = append([]string{item}, d.items...)
}

func (d *Deque) RemoveFront() string {
	if len(d.items) == 0 {
		return "" // или можно выбросить ошибку
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *Deque) RemoveRear() string {
	if len(d.items) == 0 {
		return "" // или можно выбросить ошибку
	}
	item := d.items[len(d.items)-1]
	d.items = d.items[:len(d.items)-1]
	return item
}

func (d *Deque) Size() int {
	return len(d.items)
}

func main() {
	s := "(())()())"
	q := &Deque{}
	for _, br := range s {
		if br == '(' {
			q.AddRear(string(br))
		} else {
			if q.IsEmpty() {
				panic("Ошибка")
			}
			q.RemoveRear()
		}
	}
	fmt.Print("Ок")

}
