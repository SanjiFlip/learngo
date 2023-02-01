package main

import (
	"fmt"
	"time"
)

func main() {
	round := 0
	for {
		time.Sleep(1 * time.Second)
		round++
		if round == 5 {
			continue

		}
		fmt.Println(round)
		if round > 10 {
			break
		}
	}
}
