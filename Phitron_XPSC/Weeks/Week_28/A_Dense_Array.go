package main

import "fmt"

func max(x, y int) (int, int) {
	if x > y {
		return y, x
	}

	return x, y
}

func main() {
	t := 0
	fmt.Scan(&t)
	for ; t > 0; t-- {
		n, ans := 0, 0
		fmt.Scan(&n)
		a := make([]int, n)
		for i := range a {
			fmt.Scan(&a[i])
		}
		for i := 0; i < n - 1; i++ {
			for s, l := max(a[i], a[i + 1]); s * 2 < l; {
				ans++
				s *= 2
			}
		}
		fmt.Println(ans)
	}
}
