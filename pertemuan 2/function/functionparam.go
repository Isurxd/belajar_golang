package main

import "fmt"

func helloParameter(FirstName string, lastName string) {
	fmt.Println("hello", FirstName, lastName)
}
func main() {
	helloParameter("John", "Doe")
}
