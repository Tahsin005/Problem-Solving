package main

import "fmt"

func main () {
	var x int
	fmt.Scan(&x)
	cnt := 0
	for x != 0 {
		cnt += x & 1
		x >>= 1
	}
	fmt.Println(cnt)
}