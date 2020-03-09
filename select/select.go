package main

import (
	"fmt"
	"sync"
	"time"
)

var loadIconsOnce sync.Once
var wg = sync.WaitGroup{}

func timeout(o chan<- int) {
	time.Sleep(1e9)
	o <- 1
}

func main() {
	ch := make(chan int, 5)

	fun := func(id int, cht <-chan int) {
		o := make(chan int, 1)
		go timeout(o)
		i := 10
		for {
			select {
			case data := <-cht:
				fmt.Println("goroutine id=", id, "cht get data:", data)
				i--
			case <-o:
				fmt.Println("goroutine id=", id, "timeout")
				i = 0
			}
			if i == 0 {
				break
			}
		}
		wg.Done()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go fun(i+1, ch)
	}

	for i := 0; i < 10; i++ {
		ch <- i + 1
	}
	wg.Wait()

}
