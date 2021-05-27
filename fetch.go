package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/cheggaaa/pb"
)

func fetch() {

	// for i := 0; i < len(media.Media); i++ {
	// 	fmt.Println(media.Media[i].Id, media.Media[i].Path)
	// }

	fmt.Println("Fetch image ...")

	count := len(media.Media)

	// create and start new bar
	bar := pb.StartNew(count)
	// bar.Prefix("Image")

	for i := 0; i < count; i++ {

		id := media.Media[i].Id
		url := media.Media[i].Path
		ext := filepath.Ext(url)

		// don't worry about errors
		tr := &http.Transport{
			TLSClientConfig:     &tls.Config{InsecureSkipVerify: true},
			MaxIdleConns:        20,
			MaxIdleConnsPerHost: 20,
		}
		client := &http.Client{Transport: tr}
		response, e := client.Get(url)
		if e != nil {
			log.Fatal(e)
		}
		defer response.Body.Close()

		//open a file for writing
		file, err := os.Create(folder_in + id + ext)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Use io.Copy to just dump the response body to the file. This supports huge files
		_, err = io.Copy(file, response.Body)
		if err != nil {
			log.Fatal(err)
		}

		// if i%10 == 0 {
		// 	time.Sleep(1 * time.Second)
		// }

		bar.Increment()
	}
	bar.FinishPrint("Done")
}
