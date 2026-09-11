package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	sum := 0
	for i := 1; i <= n; i++ {
		sum += (n - i) * i + 1
	}

	fmt.Print(sum)
}
