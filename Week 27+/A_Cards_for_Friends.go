package main

import "fmt"

func main () {
	var t int
	fmt.Scan(&t)
	for range t {
		var w, h, n int
		fmt.Scan(&w, &h, &n)

		cnt := 1
		for w % 2 == 0 {
			w /= 2
			cnt *= 2
		}

		for h % 2 == 0 {
			h /= 2
			cnt *= 2
		}

		if cnt >= n {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}