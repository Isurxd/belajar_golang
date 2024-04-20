package main

import "fmt"

type Adderss struct {
	City     string
	Province string
	Country  string
}

func main() {
	address1 := Adderss{"Bandung", "Jawa Barat", "Indonesia"}
	address2 := address1

	address2.City = "Jakarta"

	fmt.Println(address1)
	fmt.Println(address2)
}
