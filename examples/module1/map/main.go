package main

import (
	"fmt"
)

func main() {
	myMap := make(map[string]string, 10)
	myMap["a"] = "b"
        myMap["c"] = "d"
        myMap["e"] = "f"
	myFuncMap := map[string]func() int{
		"funcA": func() int { return 1 },
	}
	fmt.Println(myFuncMap)
	f := myFuncMap["funcA"]
	fmt.Println(f())
	value, exists := myMap["c"]
	if exists {
		println(value)
	}
	for k, v := range myMap {
		println(k, v)
	}
}
