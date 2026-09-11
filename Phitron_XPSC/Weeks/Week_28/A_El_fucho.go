package main

import "fmt"

func main() {
	var t, n int
	fmt.Scanf("%d\n", &t)
	for range t {
		fmt.Scanf("%d\n", &n)
		w := n
		l := 0
		ans := 0
		for w > 1 || l > 1 {
			ans += w / 2 + l / 2
			lose := w / 2
			w = w / 2 + w % 2
			l = l / 2 + l % 2 + lose
		}
		fmt.Println(ans + 1)
	}
}
