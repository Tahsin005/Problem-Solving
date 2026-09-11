package main

import "fmt"

func main() {
	n, m := 0, 0
	fmt.Scanf("%d %d\n", &n, &m)
	r := 0
	for a := 0; a < 1001; a++ {
		for b := 0; b < 1001; b++ {
			if a * a + b == n && a + b * b == m {
				r++
			}
		}
	}
	fmt.Println(r)
}
