package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func summ() {

}

func main() {
	r := bufio.NewReader(os.Stdin)
	num, _ := r.ReadString('\n')
	n, _ := strconv.Atoi(num)
	mas := make([]string, n)
	for i := 0; i < n; i++ {
		l := bufio.NewReader(os.Stdin)
		line, _ := l.ReadString('\n')
		mas = append(mas, line)
	}
	fmt.Println(mas)
}
