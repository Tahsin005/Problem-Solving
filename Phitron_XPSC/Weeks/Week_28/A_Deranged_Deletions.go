package main

import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := range n {
		fmt.Scan(&arr[i])
	}
	for i := 0; i + 1 < n; i++ {
		if arr[i] <= arr[i + 1] {
			continue
		}
		fmt.Println("YES")
		fmt.Println(2)
		fmt.Println(arr[i], arr[i + 1])
		return
	}
	fmt.Println("NO")
	return
}

func main() {
	var t int
	fmt.Scan(&t)
	for range t {
		solve()
	}
}
