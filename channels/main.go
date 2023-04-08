package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")
	myChan := make(chan int,2)
	wg := &sync.WaitGroup{}
	// fmt.Println(<-myChan)
	// myChan<-5
	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-myChan)
		//fmt.Println(<-myChan)

		wg.Done()
	}(myChan, wg)
	go func(ch chan int, wg *sync.WaitGroup) {
		myChan <-5
		myChan<-78
		wg.Done()
	}(myChan, wg)
	wg.Wait()

}
