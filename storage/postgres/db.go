package postgres

import (
	"fmt"
	"log"
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
)

func DB() *sqlx.DB {
	var psqlUrl = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		"localhost",
		"5432",
		"ahrorbek",
		"ahrorbek",
		"n9",
	)

	psqlConn, err := sqlx.Connect("postgres", psqlUrl)
	if err != nil {
		log.Fatal("failed to connect database: ", err)
	}

	return psqlConn
}
