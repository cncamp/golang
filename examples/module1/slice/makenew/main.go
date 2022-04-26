package main

import (
	"fmt"
	"reflect"
)

func main() {
	mySlice1 := new([]int)
	mySlice2 := make([]int, 0)
	mySlice3 := make([]int, 10)
	mySlice4 := make([]int, 10, 20)
	fmt.Printf("mySlice1 %+v\n", reflect.TypeOf(mySlice1))
	fmt.Printf("mySlice2 %+v\n", reflect.TypeOf(mySlice2))
	fmt.Printf("mySlice3 %+v\n", mySlice3)
	fmt.Printf("mySlice4 %+v\n", mySlice4)
}
