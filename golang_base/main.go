package main

import (
	"fmt"
	"time"
)

func main() {
	//两个线程轮流打印
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		for {
			ch1 <- 1
			s := <-ch2
			fmt.Println("thread1", s)

		}
	}()
	go func() {
		for {
			s := <-ch1
			fmt.Println("thread2", s)
			ch2 <- 1
		}
	}()
	time.Sleep(time.Minute)
}
