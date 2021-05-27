package main

import (
	"fmt"
	"os"
	"time"

	"github.com/common-nighthawk/go-figure"
	"github.com/manifoldco/promptui"
)

func screen() {
START:
	// clear screen
	CallClear()

	figure.NewColorFigure("DNGFOG Media", "digital", "green", true).Print()
	fmt.Println("")

	prompt := promptui.Select{
		Label: "What do you want to do",
		Items: []string{
			"Fetch images from Server (use JSON)",
			"Resize images",
			"Upload to s3 bucket",
			"Cleanup input folder",
			"Cleanup output folder",
			"Exit",
		},
		Size: 6,
	}

	key, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch key {
	case 0:
		if loadjson(JSON_FILE) {
			fmt.Println(len(media.Media))
			fetch()
			time.Sleep(1 * time.Second)
		}
		goto START
	case 1:
		resize()
		time.Sleep(3 * time.Second)
		goto START
	case 2:
		upload()
		time.Sleep(1 * time.Second)
		goto START
	case 3:
		cleanup(FOLDER_IN)
		time.Sleep(1 * time.Second)
		goto START
	case 4:
		cleanup(FOLDER_OUT)
		time.Sleep(1 * time.Second)
		goto START
	case 5:
		os.Exit(0)
	}
}
