package main

import "fmt"

func main() {
	var a1 = [1]float64{}
	fmt.Println("%#v\n", a1)
	var a2 = [2]int{-10, 1}
	fmt.Println("%#v\n", a2)
	var a3 = [3]string{"john", "paul", "George"}
	fmt.Println("%#v\n", a3)
	var a4 [4]string
	a4[0] = "x"
	a4[1] = "y"
	fmt.Println("%#v\n", a4)
}
