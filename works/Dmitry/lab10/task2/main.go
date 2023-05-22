package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var n int                  // объявляем n - количество столбцов и строк матрицы
	fmt.Scan(&n)               // ввод с клавиатуры
	matrix := make([][]int, n) // создаем матрицу размером n
	for i := 0; i < n; i++ {   // заполняем матрицу: проходимся по всем строкам
		matrix[i] = make([]int, n) // каждая строка - массив чисел, длина = n
		for j := 0; j < n; j++ {   // проходимся по каждому значению
			matrix[i][j] = rand.Intn(2) // значение = случайное число < 2
		}
	}
	fmt.Println("Матрица:")  // выводим матрицу
	for i := 0; i < n; i++ { // проходимся по каждой строке
		for j := 0; j < n; j++ { // проходимся по каждому значению
			fmt.Printf("%v\t", matrix[i][j]) // выводим значение с отступом
		}
		fmt.Println() // когда вывели значения в строке, переводим строку
	}
	symmetric := true // объявили переменную симметрии
	// симметрия относительно главной диагонали
	for i := 0; i < n; i++ { // проходимся по каждой строке
		for j := 0; j < n; j++ { // те элементы столбцов, которые выше текущей строки
			if matrix[i][j] != matrix[j][i] { //
				symmetric = false
				break
			}
		}
	} // если все элементы, расположенные выше главной диагонали равны элементам,
	// расположенным ниже главной диагонали, то матрица симметрична

	// относительно побочной диагонали
	// нужно сравнить элементы матрицы, отраженной относительно побочной диагонали,
	// с элементами исходной матрицы
	// если все элементы слева от главной диагонали равны соответствующим элементам
	// справа от главной диагонали, то матрица симметрична
	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if matrix[i][j] != matrix[n-1-j][n-1-i] {
				symmetric = false
				break
			}
		}
	}
	if symmetric { // symmetric = true
		fmt.Println("Матрица симметрична")
	} else {
		fmt.Println("Матрица не симметрична")
	}
}
