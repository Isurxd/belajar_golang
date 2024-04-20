package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func (u User) sayHello () {
	fmt.Println("hello " + u.Name)
}

func main() {
	var john User

	john.ID = 1
	john.Name = "john"
	john.Email = "john@plabs.id"

	john.sayHello()
}
