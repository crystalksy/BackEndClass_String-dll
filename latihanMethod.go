package main

import (
	"fmt"
)

type Employee struct {
	FirstName, LastName string
}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

func main() {
	e := Employee{
		FirstName: "kinasih",
		LastName:  "kinkin",
	}
	fmt.Println(e.fullName())
}
