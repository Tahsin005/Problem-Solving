package main
import (
	"fmt"
	"math"
)

func solve() {
	var s float64
	fmt.Scan(&s)
	var x float64
	x = math.Sqrt((s))
	var y float64
	y = math.Round(x)
	if float64(x - y) == 0 {
		fmt.Println(0, x)
	} else {
		fmt.Println(-1)
	}

}
func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		solve()
	}
}
