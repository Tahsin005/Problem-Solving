package main

import "fmt"

func main() {
	var tt, x, y int
	fmt.Scan(&tt)
	for tt > 0 {
		fmt.Scan(&x, &y)
		if x * 2 <= y {
			fmt.Printf("%d %d\n", x, x * 2)
		} else {
			fmt.Println("-1 -1")
		}
		tt--
	}
}