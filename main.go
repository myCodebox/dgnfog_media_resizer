package main

import (
	"bufio"
	"fmt"
	"image"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/adampresley/sigint"
	"github.com/cheggaaa/pb"
	"github.com/disintegration/imaging"
	"github.com/joho/godotenv"
)

var folder_in string
var folder_out string

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	folder_in = os.Getenv("FOLDER_IN")
	folder_out = os.Getenv("FOLDER_OUT")

	sigint.ListenForSIGINT(func() {
		fmt.Println("")
		fmt.Println("")
		fmt.Println("----------------------------------------")
		fmt.Println("Why do you press Ctrl + C?")
		fmt.Println("----------------------------------------")
		os.Exit(0)
	})

START:
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Press \"y\" to start resizing: ")

	char, _, err := reader.ReadRune()
	if err != nil {
		fmt.Println(err)
	}

	if char == 'y' || char == 'Y' {
		start()
	} else {
		goto START
	}

}

func start() {
	// read all files in folder
	files, err := os.ReadDir(folder_in)
	if err != nil {
		log.Fatal(err)
	}

	// count all files and multiply them for 4 versions
	count := len(files) * 4

	// create and start new bar
	bar := pb.StartNew(count)
	bar.Prefix("Media")

	// start with the resize
	for _, f := range files {
		normal(f.Name(), 0, 53, "_small", bar)
		normal(f.Name(), 0, 512, "", bar)
		marked(f.Name(), 0, 53, "_small_mark", "small.png", false, bar)
		marked(f.Name(), 0, 512, "_mark", "normal.png", true, bar)
	}

	// done
	bar.FinishPrint("")
	bar.FinishPrint("----------------------------------------")
	bar.FinishPrint("Yeah, done! \\o/")
	bar.FinishPrint(strconv.Itoa(count) + " Images are stored in folder \"" + folder_out + "\"")
	bar.FinishPrint("----------------------------------------")
}

func normal(name string, w int, h int, save string, bar *pb.ProgressBar) {
	src, err := imaging.Open(folder_in + name)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	rw, rh := getMaxSize(src, w, h)
	dst := imaging.Resize(src, rw, rh, imaging.Linear)

	err = imaging.Save(dst, folder_out+setName(name, save))
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}

	bar.Increment()
}

func marked(name string, w int, h int, save string, mark string, gray bool, bar *pb.ProgressBar) {
	src, err := imaging.Open(folder_in + name)
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

	err = imaging.Save(dst, folder_out+setName(name, save))
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
