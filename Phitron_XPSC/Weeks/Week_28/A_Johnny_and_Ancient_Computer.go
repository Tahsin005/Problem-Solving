package main

import "fmt"


func main() {
	var tc int
	fmt.Scan(&tc)

	for i := 0; i < tc; i++ {
		var a, b int64
		fmt.Scan(&a, &b)

		nOperation := 0
		if a > b {
			a, b = b, a
		}

		for a < b {
			if a * 8 <= b {
				a *= 8
			} else if a * 4 <= b {
				a *= 4
			} else if a * 2 <= b {
				a *= 2
			} else {
				nOperation = -1
				break
			}
			nOperation++
		}

		fmt.Println(nOperation)
	}
}
