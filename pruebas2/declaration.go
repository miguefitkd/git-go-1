package prueba2

type superheroe interface {
	getListaDeEnemigos() ([]enemigo, error)
	pelearContra(enemigo *enemigo) error
}

type personaje struct {
	nombre          string
	caracteristicas string
	peleasContra    map[string]int
}

type enemigo struct {
	personaje
}

type heroe struct {
	personaje
}

func daleNomas(heroe *superheroe) error {
	enemigos, _ := (*heroe).getListaDeEnemigos()
	for _, enemigo := range enemigos {
		_ = (*heroe).pelearContra(&enemigo)
	}
	return nil
}
