package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, delX float64 //объявляем переменные x, y, дельта x
	x = -3.5
	y = 1.6
	delX = 0.7
	for i := x; i <= 4; i += delX { // цикл: теперь i = начальному x; пока i <= 4; i += дельта x
		if i < -1.5 { // условие : если i < -1.5
			fmt.Println(i + math.Sin(y)) //вывод результата с новой строки
		} else if i >= -1.5 && i <= 2 { // если i в пределе [-1.5; 2]
			fmt.Println(y - math.Cos(i))
		} else if i > 2 { // если i > 2
			fmt.Println(math.Tan(i) + 4*y)
		}
	}
}
