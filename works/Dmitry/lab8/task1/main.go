package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, res float64 // объявили переменные x, y, res - результат
	fmt.Scan(&x, &y)      // ввод x, y с клавиатуры
	if x-y > 0 {          // условие : если x - y > 0
		res = x + math.Sin(y) // присваиваем это значение res
	} else if x-y < 0 { // иначе если x - y < 0
		res = y - math.Cos(x)
	} else if x == y { // иначе если равны
		res = math.Tan(x) + 4*y
	}
	fmt.Print(res) //выводим результат
}
