package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type studentRepo struct {
	db *sqlx.DB
}

func NewStudentRepo(db *sqlx.DB) *studentRepo {
	return &studentRepo{
		db: db,
	}
}

func (a *studentRepo) CreateStudent(ctx context.Context) error {
	return nil
}
