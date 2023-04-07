package main

import "fmt"


func main() {
	fmt.Println("welcome to a class on pointers")
	//pointers is used to pass exacte value to the target
	var ptr *int
	fmt.Println(ptr)//<nil>
	fmt.Println(&ptr)//address of the pointer
	num:=23
	var pt =&num   //& is reference it passes the address or num  now the num and pt are in same address
	fmt.Println(pt,&num)//it prints address of num and pt these are same becasue the num address is assigned to the pt
	fmt.Println(*pt)  //it prints the number 23
	*pt=*pt*3      //the value of pt is multiplied by 3 = is 69
	fmt.Println(*pt)   //it gives the out put of 69 
	fmt.Println(num)  //it gives the out put of 69 because *pt uses this address so those addresses are same result is same
	

}