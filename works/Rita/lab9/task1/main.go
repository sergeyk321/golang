package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, res float64
	y = -2                         //постоянное число
	for x = -3; x <= 3; x += 0.3 { //цикл: начинаем с х = -3, до тех пор пока х <= 3, прибавляем дельта икс = 0.3 :
		if x < -1 { //если х < -1
			fmt.Println(5 - y) //печатаем это
		} else if x >= -1 && x <= -0.1 { //иначе если х в промежутке -1 <= x <= -0.1
			fmt.Println(math.Log(x * y)) //печатаем это
		} else if x > 0.1 { //иначе если х > 0.1
			res = (math.Round(math.Abs(x+y*-2) * 10)) / 10 //округляем число до 1 знака после запятой
			fmt.Println(res)                               //печатаем результат
		}
	}
}
