package main

import "fmt"

func main() {
	var t int
	_, _ = fmt.Scan(&t)

	for i := 0; i < t; i++ {
		var (
			p1 string
			p2 string
		)
		_, _ = fmt.Scan(&p1, &p2)

		m := make(map[rune]int)

		for _, w := range p1 + p2 {
			m[w]++
		}

		fmt.Println(len(m) - 1)
	}
}
