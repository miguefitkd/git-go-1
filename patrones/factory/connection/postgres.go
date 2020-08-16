package connection

import (
	"database/sql"
	"time"
)

type Postgres struct {
	db *sql.DB
}

func (p *Postgres) Connect() error {
	return nil
}

func (p *Postgres) GetNow() (*time.Time, error) {
	return nil, nil
}

func (p *Postgres) Close() error {
	return nil
}
