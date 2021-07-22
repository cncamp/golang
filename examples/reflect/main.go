package main

import (
	"fmt"
	"reflect"
)

func main() {
	// basic type
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
	t := reflect.TypeOf(myMap)
	fmt.Println("type:", t)
	v := reflect.ValueOf(myMap)
	fmt.Println("value:", v)
	// struct
	myStruct := T{A: "a"}
	v1 := reflect.ValueOf(myStruct)
	for i := 0; i < v1.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, v1.Field(i))
	}
	for i := 0; i < v1.NumMethod(); i++ {
		fmt.Printf("Method %d: %v\n", i, v1.Method(i))
	}
	// 需要注意receive是struct还是指针
	result := v1.Method(0).Call(nil)
	fmt.Println("result:", result)
}

type T struct {
	A string
}

// 需要注意receive是struct还是指针
func (t T) String() string {
	return t.A + "1"
}
