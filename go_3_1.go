package main

import "fmt"

func main() {
	// var [nombre_variable_1, ... ,nombre_variable_n] tipo_de_dato
	// var x,y,z int
	nombre := "Hola conejo pepito"
	
	// No se puede volver a declarar la variable nombre... 
	// 
	// nombre := "Hola chavo!"
	// 
	// # command-line-arguments
    // .\go_3_1.go:12:9: no new variables on left side of :=
	// 
	
	fmt.Println(nombre)
}