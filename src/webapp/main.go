package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"./controller"
	"./model"
	"github.com/joho/godotenv"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln(err)
	}
	//log.Println(os.Getenv("ADMIN_PASS"))
	db := connectDB()
	defer db.Close()
	controller.Startup()
	http.ListenAndServe(":3000", nil)
}

func connectDB() *sql.DB {
	db, err := sql.Open(os.Getenv("DBDRIVER"), os.Getenv("DBURL"))
	if err != nil {
		log.Println("Unable to connect to DB")
		return nil
	} else {
		log.Println("Connected to DB")
		model.HandleDB(db)
		return db
	}
}
