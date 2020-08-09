package main

import (
    "fmt"
)

func main() {
   
   for i:=1;i<=10;i++ {
       fmt.Printf("for - Hola Mundo! %d\n", i)
   }
   
   fmt.Println("");
   
   y:=1
   for y<=10 {
       fmt.Printf("for/while - Hola Mundo! %d\n", y)
	   y++
   }
}