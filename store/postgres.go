package store

import (
	"context"
	"fmt"
	"net/url"

	"github.com/jackc/pgx/v4/pgxpool"
)

type postgres struct {
	conn *pgxpool.Pool
}

func NewPostgres(host, user, password, db string) IStore {
	dsn := url.URL{
		Scheme: "postgres",
		Host:   host,
		User:   url.UserPassword(user, password),
		Path:   db,
	}
	conn, err := newConn(dsn)
	if err != nil {
		panic(err)
	}
	return &postgres{
		conn: conn,
	}
}

func newConn(dsn url.URL) (*pgxpool.Pool, error) {
	q := dsn.Query()
	q.Add("sslmode", "disable")
	dsn.RawQuery = q.Encode()
	conn, err := pgxpool.Connect(context.Background(), dsn.String())
	if err != nil {
		return nil, fmt.Errorf("error connecting to db. %w", err)
	}

	return conn, nil
}
