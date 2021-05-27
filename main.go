package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var env_folder_in string
var env_folder_out string
var env_json_file string
var env_insecureSkipVerify bool
var env_maxIdleConns int
var env_maxIdleConnsPerHost int

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	env_folder_in = os.Getenv("FOLDER_IN")
	env_folder_out = os.Getenv("FOLDER_OUT")
	env_json_file = os.Getenv("JSON_FILE")

	if i, err := strconv.ParseBool(os.Getenv("INSECURE_SKIP_VERIFY")); err == nil {
		env_insecureSkipVerify = i
	}
	if i, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS")); err == nil {
		env_maxIdleConns = i
	}
	if i, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS_PER_HOST")); err == nil {
		env_maxIdleConnsPerHost = i
	}

	screen()
}
