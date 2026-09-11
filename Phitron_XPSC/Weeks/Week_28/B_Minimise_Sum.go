package main
import "fmt"

func solve() {
	var n int
	fmt.Scan(&n)
	var a = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i])
	}
	var ans int
	ans = min(a[0], a[1]) + a[0]
	fmt.Println(ans)
}

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}