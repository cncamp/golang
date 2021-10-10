package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type MyType struct {
	Name string `json:"name"`
	Address
}
type Address struct {
	City string `json:"city"`
}
func main() {
	mt := MyType{Name: "test",Address: Address{City: "shanghai"}}
	b, _ := json.Marshal(&mt)
	fmt.Println(string(b))
	myType := reflect.TypeOf(mt)
	name := myType.Field(0)
	tag := name.Tag.Get("json")
	println(tag)
	tb := TypeB{P2: "p2", TypeA: TypeA{P1: "p1"}}
	//可以直接访问 TypeA.P1
	println(tb.P1)
}

type TypeA struct {
	P1 string
}

type TypeB struct {
	P2 string
	TypeA
}
