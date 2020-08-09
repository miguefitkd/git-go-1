package main

import (
    "fmt"
)

func main(){
    
	// No se le pasa la capacidad
        slice1 := make([]int,3)
	fmt.Println( slice1 )
	fmt.Printf("Longitud del slice1: %d\n", len(slice1) )
	fmt.Printf("Capacidad del slice1: %d\n", cap(slice1) )
	
	fmt.Println()
	
	// Ahora le paso una capacidad
        slice2 := make([]int,3, 5)
	fmt.Println( slice2 )
	fmt.Printf("Longitud del slice2: %d\n", len(slice2) )
	fmt.Printf("Capacidad del slice2: %d\n", cap(slice2) )
	
	fmt.Println()
	
	// Agrego un elemento al slice sin que el programa falle.
	fmt.Println("Agrego un elemento al slice2 sin que el programa falle...")
	slice2 = append(slice2, 2)
	fmt.Println( slice2 )
	fmt.Printf("Longitud del slice2: %d\n", len(slice2) )
	fmt.Printf("Capacidad del slice2: %d\n", cap(slice2) )
	fmt.Println("Agrego otro elemento al slice2 sin que el programa falle...")
	slice2 = append(slice2, 5)
	fmt.Println( slice2 )
	fmt.Printf("Longitud del slice2: %d\n", len(slice2) )
	fmt.Printf("Capacidad del slice2: %d\n", cap(slice2) )
	fmt.Println("Intento agregar un elemento mas para que falle...")
	slice2 = append(slice2, 8)
	fmt.Println( slice2 )
	fmt.Printf("Longitud del slice2: %d\n", len(slice2) )
	fmt.Printf("Capacidad del slice2: %d\n", cap(slice2) )
	fmt.Println("=== Pero no fallo ===")
	fmt.Println("=== Se incrementaron la longitud y la capacidad del slice ===")
	
}
