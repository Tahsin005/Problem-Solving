package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var t int
	in := bufio.NewReader(os.Stdin)
	fmt.Fscan(in, &t)
	for i := 0; i < t; i++ {
		var n, d int
		fmt.Fscan(in, &n, &d)
		arr := make([]int, n)
		check := true
		for j := 0; j < n; j++ {
			fmt.Fscan(in, &arr[j])
			if arr[j] > d {
				check = false
			}
		}
		sort.Ints(arr)

		if check {
			fmt.Println("YES")
		} else {
			if arr[0] + arr[1] <= d {
				fmt.Println("YES")
			} else {
				fmt.Println("NO")
			}
		}
	}
}
