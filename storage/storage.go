package storage

import (
	"gin/storage/postgres"
	"gin/storage/repo"

	"github.com/jmoiron/sqlx"
)

type IStorage interface {
	Student() repo.StudentStorageI
	Course() repo.CourseStorageI
}

type storagePg struct {
	db          *sqlx.DB
	studentRepo repo.StudentStorageI
	courseRepo  repo.CourseStorageI
}

func NewStoragePg(db *sqlx.DB) *storagePg {
	return &storagePg{
		db:          db,
		courseRepo:  postgres.NewCourseRepo(db),
		studentRepo: postgres.NewStudentRepo(db),
	}
}

func (s storagePg) Student() repo.StudentStorageI {
	return s.studentRepo
}

func (s storagePg) Course() repo.CourseStorageI {
	return s.courseRepo
}
