package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"./controller"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load("./.env")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(os.Getenv("ADMIN_PASS"))
	fmt.Println()
	controller.Startup()
	http.ListenAndServe(":3000", nil)
}
