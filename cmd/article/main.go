package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	config "assignment/configs"
	db "assignment/db"
	"assignment/internal/app/handler"
)

func main() {
	config.InitViper()
	logPath := config.GetConfig("APP_LOG_PATH")
	logFile, err := os.OpenFile(logPath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	logger := log.New(logFile, "", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	mongoURI := config.GetConfig("MONGODSN")
	dbName := config.GetConfig("DBNAME")
	conn, err := db.Connect(mongoURI, dbName)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	handlerObj, err := handler.New(conn, logger)
	if err != nil {
		logger.Println(err)
		os.Exit(1)
	}
	serverPort := config.GetConfig("serverPort")
	logger.Println("server started on localhost:", serverPort)
	http.ListenAndServe(":"+serverPort, handlerObj)
}
