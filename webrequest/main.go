package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main() {
	response, err := http.Get(url)
	checkErr(err)
	//fmt.Println("response is of type : %T\n ", response)
	defer response.Body.Close()
	databytes, err := ioutil.ReadAll(response.Body)
	checkErr(err)
	//fmt.Println(databytes) //it prits the bytes of data
	fmt.Println(string(databytes))

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
