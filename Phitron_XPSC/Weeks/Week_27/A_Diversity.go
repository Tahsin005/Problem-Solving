package main

import (
	"fmt"
)

func main() {
	var s string
	var k int
	fmt.Scan(&s,&k)

	if len(s) < k {
		fmt.Println("impossible")
		return
	}

	chs := make(map[rune]int, 0)

	for _,ch := range s {
		chs[ch]++
	}
	out := k - len(chs)
	if out < 0 {
		out = 0
	}
	fmt.Println(out)
}