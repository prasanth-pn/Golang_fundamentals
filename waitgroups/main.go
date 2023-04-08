package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup //in go routiens the waint groups gives more flexiblity to programmer 
//it give high performance to the application because of concurrency
//ther is wait ,done and add pointer are used to make work the waitgroup 

func main() {
	start:=time.Now().UnixMilli()

	webistelist := []string{

		"https://google.co.in",
		"https://go.dev",
		"https://google.com",
		"https://github.com",
		"https://fb.com",
	}
	for _, site := range webistelist {
		go getStatusCode(site)
		wg.Add(1)
	}
	wg.Wait()
total:=time.Now().UnixMilli()-start
	fmt.Println(total)
}

func getStatusCode(endpoint string) {
	defer wg.Done()
	result, err := http.Get(endpoint)
	if err != nil {
		fmt.Println("oops in endpoint")
	}
	fmt.Printf(" %d status code for %s \n", result.StatusCode, endpoint)
}
