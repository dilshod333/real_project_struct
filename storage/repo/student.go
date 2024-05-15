package repo

import (
	"context"
)

type StudentStorageI interface {
	CreateStudent(ctx context.Context) error
}
