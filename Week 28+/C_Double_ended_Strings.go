package main
import "fmt"

func max(n int, m int) int {
	if (n > m) {
		return n
	}
	return m
}

func solve(a string, b string) int {

	m := [20][20]int{}

	var maxL int = 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			eq := 0
			if (a[i] == b[j]) {
				eq = 1
			}

			if (i * j == 0) {
				m[i][j] = eq
			} else {
				if (eq == 1) {
					m[i][j] = m[i - 1][j - 1] + 1
				}
			}
			maxL = max(m[i][j], maxL)
		}
	}
	return len(a) + len(b) - 2 * maxL
}

func main() {
	var t int
	var a,b string
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		fmt.Scan(&a, &b)
		fmt.Println(solve(a, b))
	}
}