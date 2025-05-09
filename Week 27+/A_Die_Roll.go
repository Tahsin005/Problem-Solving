package main

import "fmt"

func main() {
	var (
		y, w int32
	)

	ans := []string{"0/1", "1/6", "1/3", "1/2", "2/3", "5/6", "1/1"}

	fmt.Scan(&y)
	fmt.Scan(&w)
	if y > w {
		w = y
	}
	fmt.Println(ans[6 - w + 1])
}
