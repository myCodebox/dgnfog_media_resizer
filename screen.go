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

	myFigure := figure.NewFigure("DNGFOG Media", "digital", true)
	myFigure.Print()

	prompt := promptui.Select{
		Label: "What do you want to do",
		Items: []string{
			"Fetch from JSON to input folder",
			"Start resizing",
			"Cleanup input folder",
			"Cleanup output folder",
			"Exit",
		},
	}

	key, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	switch key {
	case 0:
		if loadjson(json_file) {
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
		cleanup(folder_in)
		time.Sleep(1 * time.Second)
		goto START
	case 3:
		cleanup(folder_out)
		time.Sleep(1 * time.Second)
		goto START
	case 4:
		CallClear()
		os.Exit(0)
	}
}
