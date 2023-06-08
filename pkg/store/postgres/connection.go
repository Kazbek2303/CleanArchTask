package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type DB struct {
	Pool *pgxpool.Pool
}

type DbConnParams struct {
	Host     string
	Port     int
	User     string
	Password string
	DbName   string
}

func OpenDB(dcp *DbConnParams) (*DB, error) {
	dsn := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", dcp.User, dcp.Password, dcp.Host, dcp.Port, dcp.DbName)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	DB := &DB{Pool: db}

	return DB, nil
}
