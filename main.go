package main

import (
	" hery-ciaputra/final-project-backend/db"
	" hery-ciaputra/final-project-backend/server"
	"log"
)

func main() {
	err := db.Connect()
	if err != nil {
		log.Println("Failed to connect to database: ", err)
	}

	server.Init()
}
