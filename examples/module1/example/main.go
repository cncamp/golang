package main

import "fmt"                                                                    
                                                                                
func swap(num1,num2 int) (int,int) {                                            
    return num2,num1                                                            
}                                                                               
                                                                                
func main() {                                                                   
    var num1, num2 int                                                          
    fmt.Print("Enter number1: ")                                                
    fmt.Scanf("%d", &num1)                                                      
    fmt.Print("enter number2: ")                                                
    fmt.Scanf("%d", &num2)                                                      
    num1, num2 = swap(num1,num2)                                                
    fmt.Printf("After Swap num1: %d and num2: %d\n", num1, num2)                
}  