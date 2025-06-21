package main

import (
	"fmt"
)

func main() {
    arr1 := [5]string{"I","am","stupid","and","weak"}
      for i :=0; i < len(arr1); i++ {
         if arr1[i] == "stupid" {
            arr1[i] = "smart"
           }
         if arr1[i] == "weak" {
            arr1[i] = "strong"
           }
         fmt.Println(arr1[i])
      }
      fmt.Println(arr1)
}
