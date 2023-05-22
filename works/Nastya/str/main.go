package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin) //создаем объект reader с помощью функцию bufio
	str, _ := reader.ReadString('\n')   //считываем строку целиком
	str = strings.TrimSuffix(str, "\n") //удаляем из строки лишний символ перевода строки (\n)
	strMas := strings.Split(str, " ")   //разбиваем строку на массив слов с помощью разделителя - пробел
	last := strMas[len(strMas)-1]       //записываем последнее слово
	last = strings.TrimSpace(last)      //удаляем из последнего слова лишние пробелы и символ перевода строки в конце
	for _, word := range strMas {       //проходим по значениям (словам) в массиве
		if last != strings.TrimSpace(word) { //если последнее слово не совпадает с тем, что есть в массиве, то печатаем его с новой строки
			fmt.Println(word)
		}
	}
}
