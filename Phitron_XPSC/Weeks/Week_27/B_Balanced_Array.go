package main

import (
	. "fmt"
)

func main() {
	var k, n int
	Scan(&k)
	for i := 0; i < k; i++ {
		Scan(&n)

		if (n / 2) % 2 == 1 {
			Println("NO")
			continue
		}
		Println("YES")

		for f, j := 0, 2; f < n / 2; f++ {
			Print(j, " ")
			j += 2
		}

		j := 1
		for f := n / 2; f < n - 1; f++ {
			Print(j, " ")
			j += 2
		}
		Println(j + n / 2)
	}
}
