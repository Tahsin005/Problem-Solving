// hey project main.go
package main

import (
    "fmt"
)

func main() {
    var x int
    fmt.Scan(&x)
    q := make([]int, x)
    m := [3]int{0, 0, 0}
    for i := 0; i < x; i++ {
        fmt.Scan(&q[i])
    }
    for i := 0; i < x; i++ {
        if q[i] == 50 {
            if m[0] == 0 {
                fmt.Printf("NO")
                return
            } else {
                m[0]--
                m[1]++
            }
        }
        if q[i] == 100 {
            if m[0] > 0 && m[1] > 0 {
                m[0]--
                m[1]--
            } else if m[0] >= 3 {
                m[0] -= 3
            } else {
                fmt.Printf("NO")
                return
            }
        }
        if q[i] == 25 {
            m[0]++
        }
    }
    fmt.Printf("YES")
}
