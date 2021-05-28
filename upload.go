package main

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/cheggaaa/pb"
)

func upload() {
	// Create a single AWS session (we can re use this if we're uploading many files)
	s, err := session.NewSession(&aws.Config{
		Region:           aws.String(S3_REGION),
		Credentials:      credentials.NewStaticCredentials(S3_AKID, S3_SECRET_KEY, S3_TOKEN),
		Endpoint:         aws.String(S3_ENDPOINT),
		DisableSSL:       aws.Bool(true),
		S3ForcePathStyle: aws.Bool(true),
	})
	if err != nil {
		log.Fatal(err)
	}

	files, err := os.ReadDir(FOLDER_OUT)
	if err != nil {
		log.Fatal(err)
	}

	// count all files and multiply them for 4 versions
	var count = 0
	for _, f := range files {
		if _, ok := FILE_TYPE[filepath.Ext(FOLDER_OUT+f.Name())]; ok {
			count++
		}
	}

	if count == 0 {
		fmt.Println("Sorry the folder \"", FOLDER_OUT, "\" is empty")
		return
	}

	fmt.Println("Resize images ...")

	// create and start new bar
	bar := pb.StartNew(count)

	// start with the resize
	for _, f := range files {
		if _, ok := FILE_TYPE[filepath.Ext(FOLDER_OUT+f.Name())]; ok {
			err = AddFileToS3(s, FOLDER_OUT+f.Name(), f.Name())
			if err != nil {
				log.Fatal(err)
			}
			bar.Increment()
		}
	}

	// done
	bar.FinishPrint("")
	bar.FinishPrint("----------------------------------------")
	bar.FinishPrint("Yeah, done! \\o/")
	bar.FinishPrint(strconv.Itoa(count) + " Images are uploaded from the folder")
	bar.FinishPrint("\"" + FOLDER_OUT + "\"  to the s3 bucket")
	bar.FinishPrint("----------------------------------------")

}

// AddFileToS3 will upload a single file to S3, it will require a pre-built aws session
// and will set file info like content type and encryption on the uploaded file.
func AddFileToS3(s *session.Session, fileDir string, fileName string) error {

	// Open the file for use
	file, err := os.Open(fileDir)
	if err != nil {
		return err
	}
	defer file.Close()

	// Get file size and read the file content into a buffer
	fileInfo, _ := file.Stat()
	var size int64 = fileInfo.Size()
	buffer := make([]byte, size)
	file.Read(buffer)

	// Config settings: this is where you choose the bucket, filename, content-type etc.
	// of the file you're uploading.
	_, err = s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(S3_BUCKET),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("private"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(size),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
	})

	return err
}
