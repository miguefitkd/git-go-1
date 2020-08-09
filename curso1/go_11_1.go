package main

import (
    "fmt"
)

func main(){
    slice1 := []int{1,2,3,4}
	slice2 := make([]int,4)
	
	fmt.Println("Valores iniciales del slice1 y el slice2...")
	fmt.Print("slice1: ")
	fmt.Println(slice1)
	fmt.Print("slice2: ")
	fmt.Println(slice2)
	fmt.Println("Copiando los valores del slice1 al slice2...")
	copy(slice2, slice1)
	fmt.Print("slice1: ")
	fmt.Println(slice1)
	fmt.Print("slice2: ")
	fmt.Println(slice2)
	fmt.Println()
	
	
	// La funcion copy copia el minimo de elementos de los dos slices.
	slice3 := []int{1,2,3,4}
	slice4 := make([]int,0) // slice con array de logitud cero
	fmt.Println("Valores iniciales del slice3 y el slice4...")
	fmt.Print("slice3: ")
	fmt.Println(slice3)
	fmt.Print("slice4: ")
	fmt.Println(slice4)
	fmt.Println("Copiando los valores del slice3 al slice4...")
	copy(slice4, slice3) // No copia elementos porque el destino tiene longitud cero
	fmt.Print("slice3: ")
	fmt.Println(slice3)
	fmt.Print("slice4: ")
	fmt.Println(slice4)
	fmt.Println()
	
	
	// La funcion copy copia el minimo de elementos de los dos slices.
	slice5 := []int{1,2,3,4}
	slice6 := make([]int,3) // slice con array de logitud 3
	fmt.Println("Valores iniciales del slice5 y el slice6...")
	fmt.Print("slice5: ")
	fmt.Println(slice5)
	fmt.Print("slice6: ")
	fmt.Println(slice6)
	fmt.Println("Copiando los valores del slice3 al slice4...")
	copy(slice6, slice5) // Copia 3 elementos porque el destino tiene longitud 3
	fmt.Print("slice5: ")
	fmt.Println(slice5)
	fmt.Print("slice6: ")
	fmt.Println(slice6)
	fmt.Println()
	
	
	// Truco...
	slice7 := []int{1,2,3,4}
	slice8 := make([]int, len(slice7))
	fmt.Println("Valores iniciales del slice7 y el slice8...")
	fmt.Print("slice7: ")
	fmt.Println(slice7)
	fmt.Print("slice8: ")
	fmt.Println(slice8)
	fmt.Println("Copiando los valores del slice7 al slice8...")
	copy(slice8, slice7)
	fmt.Print("slice7: ")
	fmt.Println(slice7)
	fmt.Print("slice8: ")
	fmt.Println(slice8)
	fmt.Println()
	
}