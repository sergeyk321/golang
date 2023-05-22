package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, z float64   // объявили переменные x, y, z
	fmt.Scan(&x, &y, &z)  // ввод с клавиатуры
	if x >= y && y >= z { // условие : если x >= y и y >= z (следовательно x >= y >= z)
		fmt.Printf("x = %v, y = %v, z = %v", x*2, y*2, z*2) //выводим удвоенные элементы с помощью форматированной строки
	} else { // иначе
		fmt.Printf("x = %v, y = %v, z = %v", math.Abs(x), math.Abs(y), math.Abs(z)) //выводим абсолютное значение
	}
}
