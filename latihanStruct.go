package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {

	//cara 1
	var Person1 Student
	Person1.FirstName = "Saddam"
	Person1.LastName = "Zericco"
	Person1.Age = 15

	fmt.Println(Person1)

	//cara 2
	var Person2 = Student{
		FirstName: "Ketut",
		LastName:  "Bambang",
		Age:       47,
	}
	fmt.Println(Person2)

	//cara 3
	Person3 := Student{"Purnamawaty", "", 44}
	fmt.Println(Person3)
}
