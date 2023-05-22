package main

import (
	"fmt"
	"math"
)

func main() {
	var e, t float64                          // объявили переменные е - точность, t - значение ряда
	fmt.Scan(&e)                              // ввод е с клавиатуры
	sum := 1.0                                // сумма
	kol := 0                                  // количество повторений
	pSum := math.Inf(-1)                      // объявили предыдущую сумму для проверки точности
	for i := 1; math.Abs(sum-pSum) > e; i++ { // цикл : начинаем с единицы, пока разность суммы предыдущей суммы больше точности е
		pSum = sum                   // обновляем предыдущую сумму
		t = 4.0 / (4*float64(i) + 3) // вычисляем значение ряда
		if i%2 == 0 {                // если элемент стоит на четной позиции в ряду
			sum += t // прибавляем его
		} else { // иначе
			sum -= t // вычитаем
		}
		kol++ // прибавляем количество повторений
	}
	fmt.Println("Сумма ряда = ", sum) // вывод суммы и количества повторенй
	fmt.Println("Количество повторений = ", kol)
}
