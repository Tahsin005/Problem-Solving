package main

import "fmt"

func main () {
	var t, n int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		arr := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&arr[j])
		}

		cnt, num := 0, -1
		for j := 0; j < n; j++ {
			if (arr[j] - num > 1) {
				cnt++
				num = arr[j]
			}
		}

		fmt.Println(cnt)
	}
}