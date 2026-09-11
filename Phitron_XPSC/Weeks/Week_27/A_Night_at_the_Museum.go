package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	last := 'a'
	cnt := int32(0)
	for _, c := range s {
		w1 := (last - c + 26) % 26
		w2 := (c - last + 26) % 26
		if w1 < w2 {
			cnt += w1
		} else {
			cnt += w2
		}
		last = c
	}

	fmt.Println(cnt)
}
