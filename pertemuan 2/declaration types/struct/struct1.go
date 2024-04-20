package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func main() {
	var john User

	john.ID = 1
	john.Name = "john"
	john.Email = "john@plabs.id"

	fmt.Printf("%#v\n", john)
}
