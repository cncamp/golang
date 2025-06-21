package main

import (
    "fmt"
)

func main() {
	DoOperation(1,decrease)
        DoOperation(1,increase)
}

func increase(a, b int) int {
	fmt.Printf("increase result is:", a+b)
        return a+b
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) int {
	fmt.Printf("decrease result is:", a-b)
        return a-b
}
