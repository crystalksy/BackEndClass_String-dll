package main

import "fmt"

func main() {
	var name string = "john"
	var nameAddress *string = &name

	fmt.Println("name(value) : ", name)
	fmt.Println("name (address) : ", &name)
	fmt.Println("nameAddress(value) : ", *nameAddress)
	fmt.Println("nameAddress (address) : ", nameAddress)

	name = "doe"
	fmt.Println("name(value) : ", name)
	fmt.Println("name (address) : ", &name)
	fmt.Println("nameAddress(value) : ", *nameAddress)
	fmt.Println("nameAddress (address) : ", nameAddress)
}
