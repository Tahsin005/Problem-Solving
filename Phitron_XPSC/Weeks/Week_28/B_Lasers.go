package main

import "fmt"

func main(){
    var t int
    fmt.Scan(&t)
    for i :=0; i < t; i++ {
        var m, n, x, y int
        fmt.Scan(&m, &n, &x, &y)
        for j := 0; j < m; j++ {
            var a int
            fmt.Scan(&a)
        }
        for j := 0; j < n; j++ {
            var b int
            fmt.Scan(&b)
        }
        fmt.Println(m + n)
    }
}