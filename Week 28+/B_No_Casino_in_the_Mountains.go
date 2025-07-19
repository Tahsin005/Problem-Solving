package main

import "fmt"

func main() {
	var testcases int
	fmt.Scan(&testcases)
	for test := 0; test < testcases; test++ {
		var n, k int
		fmt.Scan(&n, &k)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		var ans int
		var currCnt int
		for i := 0; i < n; i++ {
			if a[i] == 0 {
				currCnt++
				if currCnt == k {
					ans++
					currCnt = 0
					i++
				}
			} else {
				currCnt = 0
			}
		}

		fmt.Println(ans)
	}
}
