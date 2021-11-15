package main

import (
	"fmt"
	"testing"
)

func TestRandInt(t *testing.T) {
	for i := 0; i<3 ; i++ {
		fmt.Println(randInt(10,2000))
	}
}
