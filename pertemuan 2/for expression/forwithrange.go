package main

import "fmt"

func main() {
	names := []string{"john", "paul", "george", "ringo"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
