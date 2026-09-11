package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	var n, a, b int
	fmt.Fscanln(in, &n)
	for i := 0; i < n; i++ {
		fmt.Fscanln(in, &a, &b)
		if a != b {
			fmt.Println("Happy Alex")
			return
		}
	}
	fmt.Println("Poor Alex")
}
