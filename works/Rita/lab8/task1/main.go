package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, ans float64 //объявили переменные, ans - ответ
	fmt.Scan(&x, &y)      //записали значения в x, y с клавиатуры
	if x > y {            //условие : если x > y
		ans = math.Pow(y, x+1) - 5*x
	} else if x < y { //условие : иначе, если x < y
		ans = 6*y*2*x + 4
	} else if x == y { //условие : x = y
		ans = x + y*math.Pi/180 //переводим угол y в радианы с помощью формулы угол в радианах = угол в градусах * пи / 180 градусов
	}
	fmt.Print(ans) //выводим ответ
}
