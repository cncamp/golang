package main

import "fmt"

func main() {

	fmt.Println("this is a panic test")

	defer func() {
		fmt.Println("defer func is called")
		if err := recover(); err != nil {
			fmt.Println(err)
		}

		fmt.Println("recover main")
	}()

	panic("a panic is triggered")
}
