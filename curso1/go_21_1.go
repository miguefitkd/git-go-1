package curso1

import (
	"bufio"
	"fmt"
	"os"
)

// ExecuteGoCap21Ej1 hace tal cosa.
func ExecuteGoCap21Ej1() {
	fmt.Printf("\n\nExecuteGoCap21Ej1 <=== ****\n")
	ejecucion := readFile1("./test_1.csv")

	if ejecucion == true {
		fmt.Println("El archivo se leyo con normalidad")
	} else {
		fmt.Println("El archivo NO se pudo leer con normalidad")
	}
}

func readFile1(pathFile string) bool {
	// Abrir un archivo
	archivo, err := os.Open(pathFile)

	// HACK.-> archivo.close() se va a ejecutar sin importar en donde se realice el return
	// Opcion 1
	// defer archivo.Close()
	// Opcion 2
	defer func() {
		archivo.Close()
		fmt.Println("Ejecutando defer archivo.close()")
	}()

	if err != nil {
		// Se produjo un error
		fmt.Println("Se produjo un error al intentar acceder al archivo...")

		return false
	}

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

	return true
}
