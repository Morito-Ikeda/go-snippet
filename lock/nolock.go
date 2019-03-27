package main

import (
	"fmt"
	"time"
)

func main() {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
			fmt.Print("*")
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Printf("\n%d\n", counter)
}
