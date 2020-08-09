package curso1

import (
	"fmt"
	"io/ioutil"
)

// ExecuteGoCap19Ej1 hace tal cosa.
func ExecuteGoCap19Ej1() {
	fmt.Printf("\n\nExecuteGoCap19Ej1 <=== ****\n")
	// En "contenidoArchivo" esta todo el contenido del archivo
	contenidoArchivo, err := ioutil.ReadFile("./test_1.csv")

	if err != nil {
		// Se produjo un error
		fmt.Println("Se produjo un error al leer el archivo...")
	} else {
		// Convierte los bytes en un string, antes de imprimir
		fmt.Println(string(contenidoArchivo))
	}
}
