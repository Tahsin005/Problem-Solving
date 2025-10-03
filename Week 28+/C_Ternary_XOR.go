package main

import (
	"bufio"
	"fmt"
	"os"
)

func solve(n int, x string) string {
	result := []string{"1", "1"}
	first := true
	for i := 1; i < n; i++ {
		if x[i] == '0' {
			result[0] += "0"
			result[1] += "0"
		} else if x[i] == '2' {
			if first {
				result[0] += "1"
				result[1] += "1"
			} else {
				result[0] += "0"
				result[1] += "2"
			}
		} else {
			if first {
				result[0] += "1"
				result[1] += "0"
				first = false
			} else {
				result[0] += "0"
				result[1] += "1"
			}
		}
	}
	return result[0] + "\n" + result[1]
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var t int
	fmt.Fscan(reader, &t)

	for i := 0; i < t; i++ {
		var n int
		var x string
		fmt.Fscan(reader, &n, &x)
		fmt.Println(solve(n, x))
	}
}