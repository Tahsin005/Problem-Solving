package main

import "fmt"

func find_liar(a []int, n int) string {
	for i := 1; i < n; i++ {
		if a[i - 1] == 0 && a[i] == 0 {
			return "YES"
		}
	}

	zero_flag := false
	for i := 0; i < n; i++ {
		if a[i] == 0 {
			zero_flag = true
			break
		}
	}
	if zero_flag == false {
		return "YES"
	}
	return "NO"
}

func main() {
	var t, n int

	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&n)
		a := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&a[j])
		}
		fmt.Println(find_liar(a, n))
	}
}