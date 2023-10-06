package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var speed = 108000
	var dist = 96300000
	// var speed, dist = 108000, 96300000

	/*
		var (
			speed = 108000
			dist = 96300000
		)
	*/

	fmt.Printf("За %v дней\n", dist/speed/24)

	var num = rand.Intn(345000000) + 56000000 // генерируется от 0 до 345 млн, к резу добавляется 56 млн
	fmt.Println(num)

	fmt.Printf("За 28 дней скорость = %v", 56000000/(28*24))
}
