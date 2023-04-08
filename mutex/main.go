package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var signals = []string{"test"}
var mut sync.Mutex
//mutex is used to lock the memory address for some time 

var wg sync.WaitGroup 
func main() {
	start := time.Now().UnixMilli()

	webistelist := []string{

		"https://google.co.in",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}
	for i, site := range webistelist {
		fmt.Print(i)
		go getStatusCode(site)
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(signals)
	total := time.Now().UnixMilli() - start
	fmt.Println(total)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops in endpoint")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf(" %d status code for %s \n", result.StatusCode, endpoint)
	}
}
