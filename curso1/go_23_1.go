package curso1

import (
	"fmt"
	"io"
	"net/http"
)

// ExecuteGoCap23Ej1 hace tal cosa.
func ExecuteGoCap23Ej1() {
	fmt.Printf("\n\nExecuteGoCap23Ej1 <=== ****\n")
	// Esta linea ya levanta un servidor en el 8080
	// Pero no hay nada, asi que respondera con un 404
	// http.ListenAndServe(":8080", nil)

	// Parte de la url
	// Handler
	http.HandleFunc(
		"/bart",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Peticion para Bart")
			io.WriteString(w, "Le arranque la cabeza a tu conejo pepito!!!")
		})

	// Parte de la url
	// Hndler
	http.HandleFunc(
		"/lisa",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("Peticion para Lisa")
			io.WriteString(w, "Ese era tu conejo pepito!!!")
		})

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

// Primer parametro: struct que permite definir como se va a responder a la peticion. Representa un stream de datos de escritura.
// Segundo parametro: puntero a la informacion de la peticion recibida
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Peticion para Homero")
	// stream de datos de escritura
	// mensaje a enviar.
	io.WriteString(w, "Molesto con una copilla?")
}
