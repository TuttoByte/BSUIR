package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
)

type PostgresDB struct {
	conn *pgx.Conn
}

func NewPostgreSQL() (*PostgresDB, error) {
	conn, err := pgx.Connect(context.Background(), os.Getenv("APP_DATABASE_URL"))
	if err != nil {
		return nil, err
	}

	return &PostgresDB{conn: conn}, nil
}

func (p *PostgresDB) AddUser(name, hash string) error {
	_, err := p.conn.Exec(context.Background(),
		"INSERT INTO clients (username, password_hash) VALUES ($1, $2)",
		name, hash)
	if err != nil {
		return err
	}
	return nil

}
func (p *PostgresDB) GetUser(name string) (string, error) {
	var hash string
	err := p.conn.QueryRow(context.Background(),
		"SELECT password_hash FROM clients WHERE username=$1", name).Scan(&hash)
	if err != nil {
		return "", err
	}
	return hash, nil
}
func (p *PostgresDB) Close() error {
	err := p.conn.Close(context.Background())
	if err != nil {
		return err
	}
	return nil
}
