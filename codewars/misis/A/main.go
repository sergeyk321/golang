package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	var data []string
	var ans []string
	fmt.Scanln(&n)
	ir := bufio.NewReader(os.Stdin)
	for i := 0; i < n; i++ {
		in, _ := ir.ReadString('\n')
		in = strings.TrimSuffix(in, "\n")
		data = append(data, in)
	}
	for _, j := range data {
		k := strings.Split(j, "")
		a, _ := strconv.Atoi(k[0])
		b, _ := strconv.Atoi(k[1])
		c, _ := strconv.Atoi(k[2])
		if a-b <= c && b-a <= c {
			ans = append(ans, "YES")
		} else {
			ans = append(ans, "NO")
		}
	}
	for _, j := range ans {
		fmt.Println(j)
	}
}
