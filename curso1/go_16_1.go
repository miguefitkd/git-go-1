package curso1

import (
	"fmt"
)

type user3 interface {
	getPermisos() int
	getNombre() string
}

type admin struct {
	nombre string
}

func (admin admin) getPermisos() int {
	return 5
}

func (admin admin) getNombre() string {
	return admin.nombre
}

// --------------------------------

type editor struct {
	nombre string
}

func (editor editor) getPermisos() int {
	return 3
}

func (editor editor) getNombre() string {
	return editor.nombre
}

// --------------------------------

func auth(usuario user3) string {
	if usuario.getPermisos() >= 5 {
		return usuario.getNombre() + " tiene permisos de Administrador"
	}
	return usuario.getNombre() + " NO tiene permisos de Administrador"
}

// ExecuteGoCap16Ej1 hace tal cosa.
func ExecuteGoCap16Ej1() {
	fmt.Printf("\n\nExecuteGoCap16Ej1 <=== ****\n")
	admin := admin{"Miguel Angel"}
	fmt.Println("Autenticando usuario: " + admin.nombre)
	// La estructura Admin esta implementando la interfaz
	fmt.Println(auth(admin))
	fmt.Println()

	editor := editor{"Micaela Angelica"}
	fmt.Println("Autenticando usuario: " + editor.nombre)
	// La estructura Editor esta implementando la interfaz
	fmt.Println(auth(editor))
	fmt.Println()

	fmt.Println("Creando e iterando un Slice de Users...")
	/*
		usuarios := []user3{
			admin{
				nombre: "Celmira del Carmen",
			},
			admin{
				nombre: "Sebastian Baudilio",
			},
			editor{
				nombre: "Federico Alejandro",
			},
		}
		for _, usuario := range usuarios {
			fmt.Println(auth(usuario))
		}
	*/
	fmt.Println()
}
