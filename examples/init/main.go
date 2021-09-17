package main

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/init/a"
	_ "github.com/cncamp/golang/examples/init/b"
)

func init() {
	fmt.Println("main init")
}
func main() {

}
