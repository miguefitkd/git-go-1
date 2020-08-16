package factory

import "github.com/miguefitkd/git_go/patrones/factory/connection"

func Factory(t int) connection.DBConnection {
	switch t {
	case 1:
		return &connection.Mysql{}
	case 2:
		return &connection.Postgres{}
	default:
		return nil
	}
}
