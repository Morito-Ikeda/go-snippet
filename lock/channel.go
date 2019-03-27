package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ch := make(chan int)

	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			ch <- 1
		}()
	}

	go func() {
		for {
			i := <-ch
			counter += i
			fmt.Print("*")
			wg.Done()
		}
	}()

	wg.Wait()
	fmt.Printf("\n%d\n", counter)
}
