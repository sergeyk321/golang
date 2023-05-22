package main

import "fmt"

func main() {
	var x1, y1, x2, y2 int       //объявление переменных
	var sum, proizv string       //переменные сумма и произведение
	fmt.Scan(&x1, &y1, &x2, &y2) //записываем значения
	vesh_sum := x1 + x2          //вещественная сумма
	comp_sum := y1 + y2          //комплексная сумма
	vesh_proizv := x1*x2 - y1*y2 //вещественное произведение
	comp_proizv := y1*x2 + y2*x1 //комплексное произведение
	if comp_sum >= 0 {           //условие - если комплексная сумма больше нуля, то:
		sum = fmt.Sprintf("%d + %di", vesh_sum, comp_sum) //1) записываем сумму со знаком плюс
	} else { //иначе
		sum = fmt.Sprintf("%d - %di", vesh_sum, -comp_sum) //2) записываем сумму со знаком минус
	}
	if comp_proizv >= 0 { //условие - если комплексное произведение больше нуля, то:
		proizv = fmt.Sprintf("%d + %di", vesh_proizv, comp_proizv) //1) записываем произведение со знаком плюс
	} else { //иначе
		proizv = fmt.Sprintf("%d - %di", vesh_proizv, -comp_proizv) //2) записываем произведение со знаком минус
	}
	fmt.Println("Сумма:", sum)           //выводим сумму
	fmt.Println("Произведение:", proizv) //выводим произведение
}
