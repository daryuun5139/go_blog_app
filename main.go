package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	// "os"

	"github.com/daryuun5139/go_blog_app/api"

	_ "github.com/go-sql-driver/mysql"
)

var (
	// dbUser     = os.Getenv("USERNAME")
	// dbPassword = os.Getenv("USERPASS")
	// dbDatabase = os.Getenv("DATABASE")
	dbUser     = "daryun"
	dbPassword = "daryunpass"
	dbDatabase = "blogdb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("fail to connect DB")
		return
	}

	if err := db.Ping(); err != nil { fmt.Println(err)
		} else {
		fmt.Println("connect to DB")
		}
	

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

