package testhelpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "456456123a"
		dbname   = "FriendManagement"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	//open db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return db
}

func ConnectDBFailed() *sql.DB {
	var (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "000000"
		dbname   = "FriendManagement"
	)
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	// open db connection
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	return db

}
