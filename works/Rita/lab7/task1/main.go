package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, ans float64                //объявили переменные, ans - ответ
	fmt.Scan(&x, &y)                     //записываем значения в x, y с клавиатуры
	log := math.Log(7*x*x) / math.Log(3) //вычисляем логарифм отдельно
	ans = (6*y)*(2*math.Pow(x, 3)) +     //наш пример, функция math.Pow - x основание, 3 - степень
		(4 * math.Pow(math.E, -x)) + log //math.E - экспонента
	fmt.Println(ans) //печатаем результат
}
