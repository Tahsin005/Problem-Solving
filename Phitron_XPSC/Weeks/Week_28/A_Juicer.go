package main

import "fmt"

func main(){
	var n, b, d int
	fmt.Scan(&n, &b, &d)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	ans, sum := 0, 0
	for i := 0; i < n; i++ {
		if a[i] <=b {
			ans += a[i]
		}
		if ans > d{
			sum++
			ans = 0
		}
	}

	fmt.Println(sum)
}
