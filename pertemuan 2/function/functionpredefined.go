package main

import "fmt"

func getName() (FirstName, Lastname string) {
	return "Ferdian", "Maulana"
}
func main() {
	x, y := getName()
	fmt.Println(x,y)
}
