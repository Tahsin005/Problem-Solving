package main

import "fmt"

func main () {
	var n int
	fmt.Scan(&n)

	teams := 0
	for i := 2; i <= n; i++ {
		if n % i == 0 {
			teams++
		}
	}

	fmt.Println(teams)
}