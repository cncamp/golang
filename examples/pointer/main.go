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
	para := ParameterStruct{Name: "aaa"}
	fmt.Println(para)
	changeParameter(&para, "bbb")
	fmt.Println(para)
	cannotChangeParameter(para, "ccc")
	fmt.Println(para)
}

type ParameterStruct struct {
	Name string
}

func changeParameter(para *ParameterStruct, value string) {
	para.Name = value
}

func cannotChangeParameter(para ParameterStruct, value string) {
	para.Name = value
}
