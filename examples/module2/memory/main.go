package main

import "fmt"

// myStruct1分配在
var myStruct1 MyStruct

func main() {
	// myStruct2 分配在栈上
	var myStruct2 MyStruct
	// i分配在栈上
	i := 1
	fmt.Println(i)
	myStruct1 := MyStruct{Name: "struct1"}
	myStruct2 = MyStruct{Name: "struct2"}
	fmt.Println(myStruct1, myStruct2)
}

type MyStruct struct {
	Name string
}