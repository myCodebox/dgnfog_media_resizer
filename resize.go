package main

import (
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/cheggaaa/pb"
	"github.com/disintegration/imaging"
)

func resize() {

	// read all files in folder
	files, err := os.ReadDir(FOLDER_IN)
	if err != nil {
		log.Fatal(err)
	}

	// count all files and multiply them for 4 versions
	var count = 0
	for _, f := range files {
		if _, ok := FILE_TYPE[filepath.Ext(FOLDER_IN+f.Name())]; ok {
			count++
		}
	}
	count = count * 4

	if count == 0 {
		fmt.Println("Sorry the folder \"", FOLDER_IN, "\" is empty")
		return
	}

	fmt.Println("Resize images ...")

	// create and start new bar
	bar := pb.StartNew(count)
	// bar.Prefix("Image")

	// start with the resize
	for _, f := range files {
		if _, ok := FILE_TYPE[filepath.Ext(FOLDER_IN+f.Name())]; ok {
			normal(f.Name(), 0, 53, "-small", bar)
			normal(f.Name(), 0, 512, "", bar)
			marked(f.Name(), 0, 53, "-small-mark", "small.png", false, bar)
			marked(f.Name(), 0, 512, "-mark", "normal.png", true, bar)
		}
	}

	// done
	bar.FinishPrint("")
	bar.FinishPrint("----------------------------------------")
	bar.FinishPrint("Yeah, done! \\o/")
	bar.FinishPrint(strconv.Itoa(count) + " Images are stored in folder \"" + FOLDER_OUT + "\"")
	bar.FinishPrint("----------------------------------------")
}

func normal(name string, w int, h int, save string, bar *pb.ProgressBar) {
	src, err := imaging.Open(FOLDER_IN + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	rw, rh := getMaxSize(src, w, h)
	dst := imaging.Resize(src, rw, rh, imaging.Linear)

	err = imaging.Save(dst, FOLDER_OUT+setName(name, save))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	bar.Increment()
}

func marked(name string, w int, h int, save string, mark string, gray bool, bar *pb.ProgressBar) {
	src, err := imaging.Open(FOLDER_IN + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	overlay, err := imaging.Open("./src/mark/" + mark)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	rw, rh := getMaxSize(src, w, h)
	dst := imaging.Resize(src, rw, rh, imaging.Linear)
	if gray {
		dst = imaging.Grayscale(dst)
		dst = imaging.AdjustContrast(dst, 20)
		dst = imaging.OverlayCenter(dst, overlay, 1.0)
	} else {
		bgW := dst.Bounds().Dx() - (overlay.Bounds().Dx() + 5)
		bgH := dst.Bounds().Dy() - (overlay.Bounds().Dy() + 5)
		dst = imaging.Overlay(dst, overlay, image.Pt(bgW, bgH), 1.0)
	}

	err = imaging.Save(dst, FOLDER_OUT+setName(name, save))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	bar.Increment()
}

func getMaxSize(src image.Image, w int, h int) (rw int, rh int) {
	rw = src.Bounds().Dx()
	if h < rw {
		rw = w
	}
	rh = src.Bounds().Dy()
	if h < rh {
		rh = h
	}

	return
}

func setName(name string, save string) string {
	return name[0:len(name)-len(filepath.Ext(name))] + save + filepath.Ext(name)
}
