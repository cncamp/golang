package main

import (
  "flag"
  "fmt"
  "os"
)

func main() {
 name := flag.String("name","world","specify the name you want to say hi")
 flag.Parse()
 fmt.Println("os args is:", os.Args)
 fmt.Println("inout parameter is:", *name)
 fullString := fmt.Sprintf("hello %s from go\n", *name)
 fmt.Println(fullString)
}
