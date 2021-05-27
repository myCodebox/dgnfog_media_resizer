package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var folder_in string
var folder_out string
var json_file string

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	folder_in = os.Getenv("FOLDER_IN")
	folder_out = os.Getenv("FOLDER_OUT")
	json_file = os.Getenv("JSON_FILE")

	screen()
}
