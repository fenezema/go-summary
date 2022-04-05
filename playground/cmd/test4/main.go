package main

import (
	"fmt"
	"time"
)

func printFabs(n int) {
	fmt.Println(n, "so fabs")
}

func main() {
	for i := 0; i < 5; i += 1 {
		go printFabs(i)
	}
	time.Sleep(time.Second * 5)
}
