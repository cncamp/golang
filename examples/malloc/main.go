package main

//#cgo LDFLAGS:
//char* allocMemory();
import "C"
import (
	"fmt"
	"time"
)

func main() {
	// only loop 10 times to avoid exhausting the host memory
	for i := 1; i <= 10; i++ {
		// allocMemory malloc and memset 100Mb per time
		fmt.Printf("Allocating %dMb memory\n", i*100)
		C.allocMemory()
		time.Sleep(time.Minute)
	}
}
