package a

import (
	"fmt"

	_ "github.com/cncamp/golang/examples/init/b"
)

func init() {
	fmt.Println("init from a")
}
