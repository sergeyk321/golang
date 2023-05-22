package main

import (
	"fmt"
	"math"
)

func main() {
	var n, r, per float64                 // объявили переменные n, r, per - периметр
	fmt.Scan(&n, &r)                      // считали n и r с клавиатуры
	per = 2 * n * r * math.Tan(math.Pi/n) // периметр правильного n угольника можно выразить следующей формулой : 2nr * tan(pi/n)
	fmt.Print(per)                        // вывод результата
}
