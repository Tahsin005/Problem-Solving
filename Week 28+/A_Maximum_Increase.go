package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	var arr []int64

	for i := 0; i < n; i++ {
		var a int64 = 0
		fmt.Scan(&a)
		arr = append(arr, a)
	}

	cnt1, cnt2 := 1, 1
	for i := 0; i < n - 1; i++ {
		if arr[i] < arr[i + 1] {
			cnt1++
			if cnt2 < cnt1 {
				cnt2 = cnt1
			}
		} else {
			cnt1 = 1
		}
	}
	fmt.Print(cnt2)
}
