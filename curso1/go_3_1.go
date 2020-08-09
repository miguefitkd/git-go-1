package curso1

import "fmt"

// ExecuteGoCap3Ej1 hace tal cosa.
func ExecuteGoCap3Ej1() {
	fmt.Printf("\n\nExecuteGoCap3Ej1() <=== ****\n")

	// var [nombre_variable_1, ... ,nombre_variable_n] tipo_de_dato
	// var x,y,z int
	nombre := "Hola conejo pepito"

	// No se puede volver a declarar la variable nombre...
	//
	// nombre := "Hola chavo!"
	//
	// # command-line-arguments
	// .\go_3_1.go:12:9: no new variables on left side of :=
	fmt.Println(nombre)
}
