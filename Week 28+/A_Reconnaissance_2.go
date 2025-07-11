package main

import (
	"fmt"
	"math"
)

func main () {
	var n int
    fmt.Scan(&n)
    a := make([]int, n)
    min, ansi, ansj := 1000, 0, 0
    for i := 0; i < n; i++ { 
		fmt.Scan(&a[i]) 
	}
    for i := 0; i < n; i++ {
        u := int(math.Abs(float64(a[i]) - float64(a[(i + 1) % n])))
        if min > u {
            min = u
            ansi = i + 1
            ansj = ((i + 1) % n) + 1
        }
    }
    fmt.Println(ansi, ansj)
}