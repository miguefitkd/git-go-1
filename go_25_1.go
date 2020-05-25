package main

import (
    "net/http"
    "encoding/json"
)

// Los atributos que comienzan con minuscula no se van a mostrar en el JSON por convension
type Curso struct {
    Title string
    NumeroVideos int
}

// La convension de GO sobre JSON me va a joder en este caso
type Materia struct {
    title string
    cantidadDeCursos int
}

type Profesor struct {
    Nombre string       `json:"nombre_profesor"`
    CursosAsignados int `json:"cantidad_cursos"`
}

type Cursos []Curso

func main() {
    // Parte de la url
    // Handler
    http.HandleFunc(
                    "/", 
                    func(w http.ResponseWriter, r *http.Request) {
                        curso := Curso{"Curso de Go", 30}
                        json.NewEncoder(w).Encode(curso)
                        
                        materia := Materia{"Lorusso", 0}
                        json.NewEncoder(w).Encode(materia)
                        
                        profesor := Profesor{"Saubidet", 2}
                        json.NewEncoder(w).Encode(profesor)
                        
                        cursos := Cursos{
                                    Curso{"Go", 21},
                                    Curso{"Java", 16},
                                    Curso{"Phyton", 33},
                        }
                        json.NewEncoder(w).Encode(cursos)
    })

    http.ListenAndServe(":8080", nil)
    
}
