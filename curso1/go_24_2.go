package curso1

import (
	"fmt"
	"net/http"
)

// ExecuteGoCap24Ej2 hace tal cosa.
func ExecuteGoCap24Ej2() {
	fmt.Printf("\n\nExecuteGoCap24Ej2 <=== ****\n")
	// Servir archivos a partir de un directorio para resguardar la seguridad del sitio y sus datos

	fileServer := http.FileServer(http.Dir("public"))

	http.Handle("/", http.StripPrefix("/", fileServer))

	http.ListenAndServe(":8080", nil)
}
