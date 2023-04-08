package main

import (
	"fmt"
	"sync"
)

var score=[]int{0}
 

func main() {
	wg:= &sync.WaitGroup{}
	mut:=&sync.Mutex{}

	fmt.Println("race condition ")
	wg.Add(3)
	go func (wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("One R")
		mut.Lock()
		score=append(score,1)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	go func (wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Two R")
		mut.Lock()
		score=append(score,2)
		mut.Unlock()
		wg.Done()
	}(wg,mut)

	go func (wg *sync.WaitGroup,m *sync.Mutex){
		fmt.Println("Three R")
		mut.Lock()
		score=append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg,mut)
	wg.Wait()
	fmt.Println(score)

}