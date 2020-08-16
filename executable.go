package main

import (
	"github.com/miguefitkd/git_go/curso1"
	"github.com/miguefitkd/git_go/patrones/builder/execute_builder"
	"github.com/miguefitkd/git_go/patrones/factory/execute_factory"
	"github.com/miguefitkd/git_go/patrones/singleton/execute_singleton"
	"github.com/miguefitkd/git_go/pruebas1"
)

func main() {
	curso1.ExecuteGoCap2Ej1()
	curso1.ExecuteGoCap3Ej1()
	curso1.ExecuteGoCap4Ej1()
	// curso1.ExecuteGoCap5Ej1()
	// curso1.ExecuteGoCap5Ej2()
	curso1.ExecuteGoCap6Ej1()
	curso1.ExecuteGoCap7Ej1()
	curso1.ExecuteGoCap8Ej1()
	curso1.ExecuteGoCap9Ej1()
	curso1.ExecuteGoCap10Ej1()
	curso1.ExecuteGoCap11Ej1()
	curso1.ExecuteGoCap12Ej1()
	curso1.ExecuteGoCap13Ej1()
	curso1.ExecuteGoCap14Ej1()
	curso1.ExecuteGoCap15Ej1()
	curso1.ExecuteGoCap16Ej1()
	// curso1.ExecuteGoCap17Ej1()
	// curso1.ExecuteGoCap18Ej1()
	curso1.ExecuteGoCap19Ej1()
	curso1.ExecuteGoCap20Ej1()
	curso1.ExecuteGoCap21Ej1()
	curso1.ExecuteGoCap22Ej1()
	// curso1.ExecuteGoCap23Ej1()
	// curso1.ExecuteGoCap24Ej1()
	// curso1.ExecuteGoCap24Ej2()
	// curso1.ExecuteGoCap25Ej1()
	// curso1.ExecuteGoCap26Ej1()

	pruebas1.ExecuteGoConEj1()
	pruebas1.ExecuteGoConEj2()

	// Patterns...
	_ = execute_singleton.ExecuteSingleton()
	_ = execute_factory.ExecuteFactory()
	_ = execute_builder.ExecuteBuilder()
}
