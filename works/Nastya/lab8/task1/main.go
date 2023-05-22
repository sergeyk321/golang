package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y, ans float64 //объявление переменных x, y, ans (ответ)
	fmt.Scan(&x, &y)      //запись x и y
	if x > y {            //если x > y
		ans = math.Asin(x / 2)
	} else if x < y { //x < y
		ans = math.Max(x+4, y)
	} else if x == y { //x = y
		ans = (x * y) / math.Pow(x+y, x*y) //x+y в степени x*y
	}
	fmt.Print(ans) //ответ
}
