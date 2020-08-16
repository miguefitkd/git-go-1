package execute_singleton

import (
	"fmt"
	"sync"

	"github.com/miguefitkd/git_go/patrones/singleton"
	"github.com/miguefitkd/git_go/patrones/singleton/client_one"
	"github.com/miguefitkd/git_go/patrones/singleton/client_two"
)

// ExecuteGoConEj1 hace tal cosa
func ExecuteSingleton() error {
	fmt.Printf("\n\nExecuteSingleton <=== ****\nIncrementamos la instancias en dos puntos diferentes y mostramos el resultado por pantalla\n")
	client_one.IncrementAge()
	client_two.IncrementAge()
	p := singleton.GetInstance()
	age := p.GetAge()
	fmt.Println(age)

	fmt.Printf("Ahora vamos a incrementar de forma concurrente utilizando la libreria sync\n")
	wg := sync.WaitGroup{}
	wg.Add(200) // 200 ejecuciones concurrentes
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p = singleton.GetInstance()
	age = p.GetAge()
	fmt.Println(age)
	return nil
}
