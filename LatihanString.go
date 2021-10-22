package main

import (
	"fmt"
	"strings"
)

const (
	str    = "something"
	substr = "abc"
)

func main() {

	//len string
	sentence := "hello"
	lenSentence := len(sentence)
	fmt.Println(lenSentence)

	// compare string
	str1 := "abc"
	str2 := "abc"
	fmt.Println(str1 == str2)

	//contain
	res := strings.Contains(str, substr)
	fmt.Println(res)

	//
	value := "cat;dog"
	substring := value[4:len(value)]
	fmt.Println(substring)

	//replace
	s := "katakatakatakatas"
	t := strings.Replace(s, "s", "l", 1)
	fmt.Println(t)

	//insert
	p := "green"
	index := 2
	q := p[:index] + "HAY" + p[index:]
	fmt.Println(q)
}
