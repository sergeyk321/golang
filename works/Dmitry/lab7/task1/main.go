package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, res float64                                                        // объявили переменные x, y, res - результат
	fmt.Scan(&x, &y)                                                             //ввод с клавиатуры
	res = math.Pow(x, 0.8) + math.Sin(y) - math.Pow(y, -3)*math.Abs(math.Cos(x)) // наша формула
	fmt.Print(res)                                                               // вывод результата
}
