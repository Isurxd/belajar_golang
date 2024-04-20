package main

import "fmt"

func getHelloworld() (string, string) {
	return "Hello", "World"
}

func main() {
	x, y := getHelloworld()
	fmt.Println(x, y)
}
