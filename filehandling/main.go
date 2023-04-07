package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to files in golang")
	content := "the content is going to use in go  file4dd"
	file, err := os.Create("./myfile.txt")
	defer file.Close()
	checkNIllErr(err)
	re, err := io.WriteString(file, content) //it returns length of th file and error
	checkNIllErr(err)
	fmt.Println("the length of the writen string", re) //prints the length of the content string
	_, _ = io.WriteString(file, "\nhi prasanth")       //adding a new line file in the same file
	readFile(file.Name())

}
func readFile(filename string) {
	fmt.Println("-------------------------------------------------------------")
	databyte, err := ioutil.ReadFile(filename)
	checkNIllErr(err)
	fmt.Println(string(databyte)) //it prints the data are in the file
}
func checkNIllErr(err error) {
	if err != nil {
		panic(err)
	}

}
