package main

import "fmt"

func main() {
	var n, t int
	fmt.Scan(&n, &t)

	if t != 10 {
		for i := 0; i < n; i++ {
			fmt.Print(t)
		}
	} else if n == 1 && t == 10 {
		fmt.Println(-1)
	} else {
		fmt.Print(1)
		for i := 0; i < n - 1; i++ {
			fmt.Print(0)
		}
		return
	}
}
