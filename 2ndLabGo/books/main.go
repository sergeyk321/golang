package main

import (
	"fmt"
	"strings"
)

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
	s := "Преступление и наказание, Три товарища, Униженные и оскорбленные, Портрет Дориана Грея, 20000 лье под водой"
	str := strings.Split(s, ", ")
	deq1 := &Deque{}      //отсортирован
	deq2 := &Deque{}      //не отсортирован
	deq1.AddFront(str[0]) //начальный

	for i := 1; i < len(str); i++ {
		if str[i] <= deq1.items[0] {
			deq1.AddFront(str[i]) //ставим первым
		} else if str[i] > deq1.items[deq1.Size()-1] {
			deq1.AddRear(str[i]) //ставим последним
		} else {
			for str[i] > deq1.items[0] { //убираем в неотсортированную
				deq2.AddFront(deq1.RemoveRear())
			}
			deq1.AddFront(str[i]) //вставляем наш элемент
			for !deq2.IsEmpty() { //возвращаем в отсортированный
				deq1.AddFront(deq2.RemoveRear())
			}
		}
	}
	for !deq1.IsEmpty() {
		fmt.Println(deq1.RemoveFront())
	}
}
