package main

import (
	"fmt"
)

func main() {
	address1 := address{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := &address1

	address2.city = "jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}
