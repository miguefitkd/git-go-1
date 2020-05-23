package main

import (
  "fmt"
  "strconv"
)

func main() {
	edad := 36
	fmt.Println(edad)
	
	// (type untyped string) as type int
	// fmt.Println("Mi edad es: " + edad)
	
	edad_str := strconv.Itoa( edad )
	fmt.Println("Mi edad es: " + edad_str)
	
	// El segundo valor es el error.
	edad_int,_ := strconv.Atoi( edad_str )
	fmt.Println(edad_int + 10)
}