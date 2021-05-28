package main

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	FOLDER_IN               string
	FOLDER_OUT              string
	JSON_FILE               string
	INSECURE_SKIP_VERIFY    bool
	MAX_IDLE_CONNS          int
	MAX_IDLE_CONNS_PER_HOST int
	S3_REGION               string
	S3_BUCKET               string
	S3_ENDPOINT             string
	S3_AKID                 string
	S3_SECRET_KEY           string
	S3_TOKEN                string

	// myFileType = map[string]int{
	FILE_TYPE = map[string]int{
		".jpg":  1,
		".jpeg": 2,
		".png":  3,
		".gif":  4,
		".tif":  5,
		".tiff": 6,
		".bmp":  7,
	}
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	FOLDER_IN = os.Getenv("FOLDER_IN")
	FOLDER_OUT = os.Getenv("FOLDER_OUT")
	JSON_FILE = os.Getenv("JSON_FILE")

	if i, err := strconv.ParseBool(os.Getenv("INSECURE_SKIP_VERIFY")); err == nil {
		INSECURE_SKIP_VERIFY = i
	}
	if i, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS")); err == nil {
		MAX_IDLE_CONNS = i
	}
	if i, err := strconv.Atoi(os.Getenv("MAX_IDLE_CONNS_PER_HOST")); err == nil {
		MAX_IDLE_CONNS_PER_HOST = i
	}

	S3_REGION = os.Getenv("S3_REGION")
	S3_BUCKET = os.Getenv("S3_BUCKET")
	S3_ENDPOINT = os.Getenv("S3_ENDPOINT")
	S3_AKID = os.Getenv("S3_AKID")
	S3_SECRET_KEY = os.Getenv("S3_SECRET_KEY")
	S3_TOKEN = os.Getenv("S3_TOKEN")

	screen()
}
