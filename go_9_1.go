package main

import (
    "fmt"
)

func main() {
    var sliceVacio []int
	fmt.Println(sliceVacio)
	if sliceVacio == nil {
	    fmt.Println("El slice esta vacio")
	} else {
	    fmt.Println("El slice NO esta vacio")
	}
	fmt.Println()
	
	slice := []int{1,2,3,4}
	fmt.Println(slice)
	if slice == nil {
	    fmt.Println("El slice esta vacio")
	} else {
	    fmt.Println("El slice NO esta vacio")
	}
	fmt.Printf("El tamanio del slice es: %d\n", len(slice))
	fmt.Println()
	
	// Estructura de un slice tiene...
	// Puntero al arreglo
	// Longitud del arreglo
	// Capacidad
	
	arreglo := [4]int {1,2,3}
	fmt.Println(arreglo)
	slice1 := arreglo[1:2] // Se toma desde la posicion 1 hasta la posicion 2 (no inclusivo)
	fmt.Println(slice1)
	slice2 := arreglo[:2]  // Se toma hasta la posicion 2 (no inclusivo)
	fmt.Println(slice2)
	slice3 := arreglo[:0]  // Vacio
	fmt.Println(slice3)
	slice4 := arreglo[0:4]  // Vacio
	fmt.Println(slice4)
	slice5 := arreglo[1:]  // Vacio
	fmt.Println(slice5)
	fmt.Println()
}