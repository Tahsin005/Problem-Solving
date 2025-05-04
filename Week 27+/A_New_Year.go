package main

import "fmt"
import "sort"

func main(){
    arr := make([]int, 3)
    fmt.Scan(&arr[0], &arr[1], &arr[2])
    sort.Ints(arr)
    fmt.Print(arr[2] - arr[0])
}
