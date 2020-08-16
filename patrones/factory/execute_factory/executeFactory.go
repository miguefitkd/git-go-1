package execute_factory

import (
	"fmt"

	"github.com/miguefitkd/git_go/patrones/factory"
)

func ExecuteFactory() error {
	fmt.Printf("\n\nExecuteFactory <=== ****\nAca deberia haber utilizado un Scanf pero que fiaca che....\n")

	var typeConnection int = 1
	conn := factory.Factory(typeConnection)
	if conn == nil {
		fmt.Printf("Type: %d - Motor no valido\n", typeConnection)
	} else {
		now, _ := conn.GetNow()
		if now == nil {
			fmt.Printf("Type: %d - No se pudo recuperar la fecha\n", typeConnection)
		}
	}

	typeConnection = 2
	conn = factory.Factory(typeConnection)
	if conn == nil {
		fmt.Printf("Type: %d - Motor no valido\n", typeConnection)
	} else {
		now, _ := conn.GetNow()
		if now == nil {
			fmt.Printf("Type: %d - No se pudo recuperar la fecha\n", typeConnection)
		}
	}

	typeConnection = 3
	conn = factory.Factory(typeConnection)
	if conn == nil {
		fmt.Printf("Type: %d - Motor no valido\n", typeConnection)
	} else {
		now, _ := conn.GetNow()
		if now == nil {
			fmt.Printf("Type: %d - No se pudo recuperar la fecha\n", typeConnection)
		}
	}

	return nil
}
