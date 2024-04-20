package main

import "fmt"

func main () {
	address1 := Address{"bandung", "jawa barat", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}