package main

import (
    "net/http"
    "fmt"
)

func main() {
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
