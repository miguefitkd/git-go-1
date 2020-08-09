package curso1

import (
	"bufio"
	"fmt"
	"os"
)

// ExecuteGoCap20Ej1 hace tal cosa.
func ExecuteGoCap20Ej1() {
	fmt.Printf("\n\nExecuteGoCap20Ej1 <=== ****\n")
	// Abrir un archivo
	archivo, err := os.Open("./test_1.csv")

	if err != nil {
		// Se produjo un error
		fmt.Println("Se produjo un error al intentar acceder al archivo...")
	} else {

		// Generar un scanner
		scanner := bufio.NewScanner(archivo)

		index := 0

		// Iterar todas las lineas hasta el final del archivo
		for scanner.Scan() {
			index++
			linea := scanner.Text()
			fmt.Printf("Linea %d: ", index)
			fmt.Println(linea)
		}
		archivo.Close()
	}
}
