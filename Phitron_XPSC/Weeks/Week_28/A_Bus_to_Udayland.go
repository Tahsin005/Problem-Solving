package main

import (
	"fmt"
	"strings"
)

func main() {
	var n int
	b := []string{}
	ok := false
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		var x, y string
		fmt.Scan(&x)
		if (!ok) {
			y = strings.Replace(x, "OO", "++", 1)
			ok = x!=y
			x = y
		}
		b = append(b, x)
	}
	if ok {
		fmt.Println("YES")
		fmt.Println(strings.Join(b, "\n"))
	} else {
		fmt.Println("NO")
	}
}