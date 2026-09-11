package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	e := make([]int, 3)
	t := 0
	for k := 0; k < n; k++ {
		fmt.Scan(&t)
		e[k % 3] += t
	}
	if e[0] > e[1] && e[0] > e[2] {
		fmt.Println("chest")
	} else if e[1] > e[2] {
		fmt.Println("biceps")
	} else {
		fmt.Println("back")
	}
}