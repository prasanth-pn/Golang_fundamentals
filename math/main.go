package main

import (
	"fmt"
	"math/big"
	"crypto/rand"
)

func main() {
	fmt.Println("welcome to the math packages")
	//randoom number generating
	// 	rand.Seed(time.Now().UnixNano())
	// 	fmt.Println(rand.Intn(3))

	// crypto random generating
	randomNumber,_ := rand.Int(rand.Reader,big.NewInt(34))
	fmt.Println(randomNumber)

}
