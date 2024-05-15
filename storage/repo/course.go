package repo

import (
	"context"
)

type CourseStorageI interface {
	CreateCourse(ctx context.Context) error
}
