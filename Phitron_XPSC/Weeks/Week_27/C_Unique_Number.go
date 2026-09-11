package main

import "fmt"

func main() {
	var tt int
	fmt.Scan(&tt)
	for i := 0; i < tt; i++ {
		var n int
		fmt.Scan(&n)

		d := 9
		ans := 0
		p := 1
		for n > 0 && d > 0 {
			if n >= d {
				n -= d
				ans += p * d
				p *= 10
			}
			d--
		}
		if n == 0 {
			fmt.Println(ans)
			continue
		}
		fmt.Println(-1)
	}
}
