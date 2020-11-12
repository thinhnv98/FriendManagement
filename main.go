package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"S3_FriendManagement_ThinhNguyen/routes"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//Load .env file
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error load .env file")
	}

	//Connect DB
	db := ConnectDB()
	defer db.Close()

	//create routes
	r := routes.CreateRoutes(db)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func ConnectDB() *sql.DB {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DBNAME")
	)

	pgsqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", pgsqlInfo)
	if err != nil {
		panic(err)
	}

	return db
}
