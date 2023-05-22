package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y float64
	y = 3                         //y всегда 3
	for x = 2; x <= 8; x += 0.8 { //цикл от 2 до 8, дельта 0.8
		if x < 3 { //если x < 3
			fmt.Println(math.Atan(x / 2)) //ответ
		} else if x >= 3 && x <= 6 { //или если 3<=x<=6
			fmt.Println(x + 4*y)
		} else if x > 6 { //или если x > 6
			fmt.Println((x * y) / math.Pow(x+y, x*y))
		}
	}
}
