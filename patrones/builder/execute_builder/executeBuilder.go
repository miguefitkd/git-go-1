package execute_builder

import (
	"fmt"

	"github.com/miguefitkd/git_go/patrones/builder"
)

func ExecuteBuilder() error {
	fmt.Printf("\n\nExecuteBuilder <=== ****\n")

	director := builder.Sender{}
	receptor := "El señor X"
	message := "La puerca esta en la pocilga"

	// En la practica definirias un builder... y crearia los mensajes de una u otra forma.

	// Primer intento con JSON
	fmt.Printf("Json Builder...\n")
	jsonBuilder := builder.JSONMessageBuilder{}
	director.SetBuilder(&jsonBuilder)
	messageRepresented, err := director.BuildMessage(receptor, message)
	fmt.Printf("Error: %s, Format: %s, Body: %s", err, messageRepresented.Format, messageRepresented.Body)

	// Segundo intento con XML
	fmt.Printf("Xml Builder...\n")
	xmlBuilder := builder.XMLMessageBuilder{}
	director.SetBuilder(&xmlBuilder)
	messageRepresented, err = director.BuildMessage(receptor, message)
	fmt.Printf("Error: %s, Format: %s, Body: %s", err, messageRepresented.Format, messageRepresented.Body)

	return nil
}
