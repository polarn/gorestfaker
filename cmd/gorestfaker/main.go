package main

import (
	"log"
	"os"
	"strconv"

	"github.com/polarn/gorestfaker/internal/httpserver"
)

var (
	ServerPortStr string
	ServerPort int = 8080
)

func init() {
	ServerPortStr = os.Getenv("SERVER_PORT")
	if len(ServerPortStr) > 0 {
		port, err := strconv.Atoi(ServerPortStr)
		if err != nil {
			log.Printf("Error converting SERVER_PORT to integer, using default. Error: %s\n", err.Error())
		} else {
			ServerPort = port
			log.Printf("Using SERVER_PORT environment variable as server port.\n")
		}
	}
}

func main() {
	log.Printf("Starting gorestfaker on port %d", ServerPort)
	server := httpserver.NewServer(ServerPort)
	log.Fatal(server.Start())
}
