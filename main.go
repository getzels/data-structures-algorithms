package main

import "fmt"

func main() {
    var numbers = []int{1,2,3,4,5}
    
    for _, number1 := range numbers {
       for _, number2 := range numbers {
       
            fmt.Println("Part ",number1,number2)        
        }
    }
}