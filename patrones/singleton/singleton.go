package singleton

import "sync"

var (
	p    *person
	once sync.Once // esta variable va a permitir ejecutar una unica vez una porcion de codigo
)

func GetInstance() *person {
	once.Do(func() {
		// Esto se ejecuta una unica vez por paquete, es decir que no podes
		// volver a utilizar el once.Do en otra funcion de este paquete, va
		// a ejecutar el primero que encuentre y al/los siguiente/s lo/s va
		// a ignorar
		p = &person{}
	})
	return p
}

type person struct {
	name string
	age  int
	mux  sync.Mutex // Este mutex va a permitir bloquear momentaneamente el proceso y desbloquearlo
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) GetAge() int {
	p.mux.Lock()
	defer p.mux.Unlock() // Utilizo defer para desbloquear luego del return
	return p.age
}

func (p *person) SetName(n string) {
	p.name = n
}

func (p *person) SetAge(a int) {
	p.mux.Lock()
	p.age = a
	p.mux.Unlock()
}

func (p *person) IncrementAge() {
	p.mux.Lock()
	p.age++
	p.mux.Unlock()
}
