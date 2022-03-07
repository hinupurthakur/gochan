package main

import (
	"fmt"
	"sync"
)

func main() {
	even, odd := make(chan bool), make(chan bool)

	wg := sync.WaitGroup{}
	go func() {
		start := 10
		for {
			start += 2
			even <- true
			fmt.Println(start)
			<-odd
			if start >= 20 {
				break
			}
		}
		close(even)
	}()

	wg.Add(1)
	go func() {
		start := 1
		for {
			odd <- true
			fmt.Println(start)
			start += 2
			_, ok := <-even
			if ok == false {
				break
			}
		}
		wg.Done()
	}()

	<-odd

	wg.Wait()
	fmt.Print("\n")
}
