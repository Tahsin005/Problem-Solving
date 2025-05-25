package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	dic := make(map[string]string)
	for i := 0; i < m; i++ {
		var a, b string
		fmt.Scan(&a, &b)
		if len(b) < len(a) {
			dic[a] = b
		} else {
			dic[a] = a
		}
	}

	for i := 0; i < n; i++ {
		var s string
		fmt.Scan(&s)
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(dic[s])
	}
	fmt.Println()
}
