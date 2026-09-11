package main

import "fmt"

func main () {
	var tt int
	fmt.Scan(&tt)

	for ii := 0; ii < tt; ii++ {
		var n int
		fmt.Scan(&n)
		t := 0
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&a[i])
			if a[i] == 1 {
				t++
			}
		}
		fmt.Println(n - t / 2)
	}
}