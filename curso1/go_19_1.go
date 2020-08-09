package main

import (
    "fmt"
	"io/ioutil"
)

func main() {
    
    // En "contenido_archivo" esta todo el contenido del archivo
	contenido_archivo,err := ioutil.ReadFile("./test_1.csv")
	
	if ( err != nil) {
	    // Se produjo un error
		fmt.Println("Se produjo un error al leer el archivo...");
	} else {
	    // Convierte los bytes en un string, antes de imprimir
	    fmt.Println( string(contenido_archivo) );
	}
	
}
