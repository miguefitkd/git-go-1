package curso1

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Los atributos que comienzan con minuscula no se van a mostrar en el JSON por convension
type curso struct {
	Title        string
	NumeroVideos int
}

// La convension de GO sobre JSON me va a joder en este caso
type materia struct {
	title            string
	cantidadDeCursos int
}

type profesor struct {
	Nombre          string `json:"nombre_profesor"`
	CursosAsignados int    `json:"cantidad_cursos"`
}

type cursos []curso

// ExecuteGoCap25Ej1 hace tal cosa.
func ExecuteGoCap25Ej1() {
	fmt.Printf("\n\nExecuteGoCap25Ej1 <=== ****\n")
	// Parte de la url
	// Handler
	http.HandleFunc(
		"/",
		func(w http.ResponseWriter, r *http.Request) {
			curso := curso{
				Title:        "Curso de Go",
				NumeroVideos: 30,
			}
			json.NewEncoder(w).Encode(curso)

			materia := materia{
				title:            "Lorusso",
				cantidadDeCursos: 0,
			}
			json.NewEncoder(w).Encode(materia)

			profesor := profesor{"Saubidet", 2}
			json.NewEncoder(w).Encode(profesor)
			/*
				cursos := cursos{
					curso{
						Title:        "Go",
						NumeroVideos: 21,
					},
					curso{
						Title:        "Java",
						NumeroVideos: 16,
					},
					curso{
						Title:        "Phyton",
						NumeroVideos: 33,
					},
				}
				json.NewEncoder(w).Encode(cursos)
			*/
		})
	http.ListenAndServe(":8080", nil)
}
