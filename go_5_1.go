package main

import (
  "fmt"
//  "strconv"
)

// Leer e imprimir datos

func main() {
	fmt.Println("Hola Mundo!")
	
	edad := 36
	fmt.Printf("Mi edad es: %d\n", edad)
	
	edad = 37
	fmt.Printf("Mi edad es: %v\n", edad)
	
	nombre := "Miguel"
	fmt.Printf("y mi nombre es: %v\n", nombre)
	nombre  = "Angel"
	fmt.Printf("y mi segundo nombre es: %s\n", nombre)
	
	flag := true
	fmt.Printf("flag_1: %v\n", flag)
	flag = false
	fmt.Printf("flag_2: %v\n", flag)
	flag = true
	fmt.Printf("flag_3: %t\n", flag)
	
	precio := 14.36
	fmt.Printf("Precio(v): %v\n", precio)
	fmt.Printf("Precio(f): %f\n", precio)
	
	// ------------------------------------
	fmt.Print("Define tu edad: ")
	fmt.Scanf("%d\n", &edad)
	fmt.Printf("La edad ingresada fue: %v\n", edad)
	
	apellido := "Fulanito"
	fmt.Print("Define un primer apellido: ")
	fmt.Scanf("%v\n", &apellido)
	fmt.Printf("El apellido ingresado fue(v): %v\n", apellido)
	fmt.Print("Define un segundo apellido: ")
	fmt.Scanf("%s\n", &apellido)
	fmt.Printf("El apellido ingresado fue(s): %s\n", apellido)
	
}