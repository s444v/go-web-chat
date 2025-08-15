package main

import (
	"log"
	"os"

	"github.com/s444v/go-web-chat/pkg/database"
	"github.com/s444v/go-web-chat/pkg/server"
)

func main() {
	logger := log.New(os.Stdout, "Info: ", log.Ldate|log.Ltime|log.Llongfile)
	err := database.DBinit()
	if err != nil {
		logger.Fatalf("Ошибка в работе с базой данных: %v", err)
	}
	defer database.DB.Close()
	router := server.NewServer()
	router.Run(":8080")
}
