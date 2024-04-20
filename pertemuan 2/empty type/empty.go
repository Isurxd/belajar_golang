package main

import "fmt"

func ups() interface{} {
	//return 1
	// return true
	return "ups"
}

func main() {
	data := ups()
	fmt.Println(data)
}
