package main

import (
	"fmt"
)

func main() {
	str := "a string value"
	pointer := &str
	anotherString := *&str
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)
	str = "changed string"
	fmt.Println(str)
	fmt.Println(pointer)
	fmt.Println(anotherString)
}
