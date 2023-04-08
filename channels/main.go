package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("channels in golang")
	myChan := make(chan int,5)
	wg := &sync.WaitGroup{}
	// fmt.Println(<-myChan)
	// myChan<-5
	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {

		val,isChanelOpen:=<-myChan
		fmt.Println(isChanelOpen)
		fmt.Println(val)

		wg.Done()
	}(myChan, wg)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		myChan <-0
		//myChan <-78
		close(myChan)
		wg.Done()
	}(myChan, wg)
	wg.Wait()

}
