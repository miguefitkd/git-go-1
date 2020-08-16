package connection

import (
	"database/sql"
	"time"
)

type Mysql struct {
	db *sql.DB
}

func (m *Mysql) Connect() error {
	return nil
}

func (m *Mysql) GetNow() (*time.Time, error) {
	return nil, nil
}

func (m *Mysql) Close() error {
	return nil
}
