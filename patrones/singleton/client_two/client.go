package client_two

import "github.com/miguefitkd/git_go/patrones/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
