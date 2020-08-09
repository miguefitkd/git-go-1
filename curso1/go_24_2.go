package main

import (
    "net/http"
)

func main() {
    // Servir archivos a partir de un directorio para resguardar la seguridad del sitio y sus datos
    
    fileServer := http.FileServer(http.Dir("public"))
    
    http.Handle("/", http.StripPrefix("/", fileServer) )
    
    http.ListenAndServe(":8080", nil)
    
}
