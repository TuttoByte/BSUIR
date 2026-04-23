package models

import (
	"context"
	"github.com/jackc/pgx/v5"
	"os"
	"reflect"
	"strings"
)

type PostgresDB[T any] struct {
	conn *pgx.Conn
}

func NewPostgresDB[T any](ctx context.Context, conn *pgx.Conn) (*PostgresDB[T], error) {
	conn, err := pgx.Connect(ctx, os.Getenv("APP_DATABASE_URL"))
	if err != nil {
		return nil, err
	}
	return &PostgresDB[T]{conn: conn}, nil
}

func (p *PostgresDB[T]) Create(ctx context.Context, item *T) error {

}
func (p *PostgresDB[T]) FindByID(ctx context.Context, id uint64) (*T, error) {

}
func (p *PostgresDB[T]) Update(ctx context.Context, item *T) error {
	vtype := reflect.TypeOf(item)
	var insertValues strings.Builder
	for field := range vtype.Fields() {
		_, err := insertValues.Write([]byte(field.Name))
		if err != nil {
			return err
		}
	}

	var values strings.Builder
	for i := range vtype.NumField() {
		vtype.
	}
}
func (p *PostgresDB[T]) Delete(ctx context.Context, id uint64) error {

}
