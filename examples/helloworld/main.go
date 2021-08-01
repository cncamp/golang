package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "world", "specify the name you want to say hi")
	flag.Parse()
	fmt.Println("os args is:", os.Args)
	fmt.Println("input parameter is:", *name)
	fmt.Printf("Hello %s from Go\n", *name)
}
