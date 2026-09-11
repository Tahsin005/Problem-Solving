package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i < t; i++ {
		var n int
		fmt.Scan(&n)
		messages := make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&messages[j])
		}
		cnt := countWalk(messages)
		if cnt >= 2 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func countWalk(messages []int) (cnt int) {
	lastMessage := 0
	n := len(messages)
	if messages[n - 1] != 1440 {
		messages = append(messages, 1440)
	}
	for _, nowMessage := range messages {
		if nowMessage - lastMessage >= 120 {
			cnt++
		}
		if nowMessage - lastMessage >= 240 {
			cnt++
		}
		lastMessage = nowMessage
	}
	return
}