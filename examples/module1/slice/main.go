package main
import (
 "fmt"
)
func main() {
  myArray := [5]int{1,2,3,4,5}
  mySlice := myArray[1:5]
  fmt.Printf("mySlice %+v\n", mySlice)
  fullSlice := myArray[:]
  fmt.Printf("fullSlice %+v\n", fullSlice)
  remove3rdItem := deleteItem(fullSlice,4)
  fmt.Printf("remove3rdItem %+v\n", remove3rdItem)
}
func deleteItem(slice []int,index int)[]int {
  return append(slice[:index],slice[index+1:]...)
}


