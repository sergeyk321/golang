package main

import (
	"fmt"
)

type Deque struct {
	items []rune
}

func (d *Deque) IsEmpty() bool {
	return len(d.items) == 0
}

func (d *Deque) AddRear(item rune) {
	d.items = append(d.items, item)
}

func (d *Deque) AddFront(item rune) {
	d.items = append([]rune{item}, d.items...)
}

func (d *Deque) RemoveFront() rune {
	if len(d.items) == 0 {
		return 0 // или можно выбросить ошибку
	}
	item := d.items[0]
	d.items = d.items[1:]
	return item
}

func (d *Deque) RemoveRear() rune {
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

func enc(r rune, deq *Deque) string {
	for i := 0; i < deq.Size(); i++ {
		t := deq.RemoveRear()
		if string(t) == string(r) {
			val1 := deq.RemoveRear()
			val2 := deq.RemoveRear()
			deq.AddRear(val1)
			deq.AddRear(val2)
			return string(val1)
		}
		deq.AddRear(t)
	}
	return string(r)
}

func main() {
	s := "абвгд"
	a := "абвгдежзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯ"

	deq := &Deque{}
	for _, r := range a {
		deq.AddRear(r)
	}

	ans := ""
	for _, r := range s {
		ans += string(enc(r, deq))
	}
	fmt.Println(ans)
}
