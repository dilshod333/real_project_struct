package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"
)

type courseRepo struct {
	db *sqlx.DB
}

func NewCourseRepo(db *sqlx.DB) *courseRepo {
	return &courseRepo{
		db: db,
	}
}

func (a *courseRepo) CreateCourse(ctx context.Context) error {
	return nil
}
