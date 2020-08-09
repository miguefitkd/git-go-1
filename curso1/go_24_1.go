package curso1

import (
	"fmt"
	"net/http"
)

// ExecuteGoCap24Ej1 hace tal cosa.
func ExecuteGoCap24Ej1() {
	fmt.Printf("\n\nExecuteGoCap24Ej1 <=== ****\n")

	// Parte de la url
	// Handler
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			fmt.Println("PATH: " + r.URL.Path)
			// http.ServeFile(w, r, "index.html")
			http.ServeFile(w, r, r.URL.Path[1:])
		})

	http.ListenAndServe(":8080", nil)
}
