package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	a, _ := r.ReadString('\n')
	strings.TrimSuffix(a, "\n")
	fmt.Println(findSymbol(a))
}

func findSymbol(a string) string {
	var (
		maxC int
		maxS string
	)
	ans := make(map[rune]int)
	for _, j := range a {
		ans[j]++
		if ans[j] > maxC {
			maxC = ans[j]
			maxS = string(j)
		}
	}
	return maxS + " повторяется " + strconv.Itoa(maxC) + " раз(а)"
}
