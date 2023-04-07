package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	ap:=User{}
	ap.Status=true
	fmt.Println("entered into methods or interface")
	fmt.Println(ap)
	ap.GetStatus()
	ap.NewMail()
	fmt.Println("the email is ",ap.Email)

}
func (u *User)GetStatus(){
	fmt.Println("u.Status",u.Status)
}
func (u User)NewMail(){
	u.Email="prasanth"
	fmt.Println("email is ",u.Email)
	
}
