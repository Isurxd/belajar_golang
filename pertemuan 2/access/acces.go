package main

import "fmt"

func main() {
	fmt.Println(controller.ReadUserList())
}

type user struct {
	id    int    `json: "id"`
	name  string `json: "name"`
	email string `json: "email"`
}

func ReadUserList() (users []user) {
	users = make([]user, 2)

	users[1] = user{
		id:    1,
		name:  "Surya Tripatih",
		email: "surya@plabs.id",
	}

	users[1] = user{
		id:    2,
		name:  "Diva Dien Alhaq",
		email: "diva@plabs.id",
	}
	return
}

