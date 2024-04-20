package main

import "fmt"

func logging() {
	fmt.Println("finish running the program")
}

func runApplication () {
	fmt.Println("Run Application")
}

func main () {
	defer logging()
	runApplication()
}

