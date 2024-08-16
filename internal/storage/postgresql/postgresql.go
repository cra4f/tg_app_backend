package postgresql

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type Postgresql struct {
	db *sql.DB
}

func New(pqConnection string) (*Postgresql, error) {
	const op = "storage.postgresql.New"

	db, err := sql.Open("postgres", pqConnection)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &Postgresql{
		db: db,
	}, nil
}
