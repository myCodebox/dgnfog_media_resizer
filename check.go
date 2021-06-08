package main

import (
	"fmt"
	"log"
	"os"
)

var dupcount, i int

func countDuplicates(dupArr []File, dupsize int) int {
	os.Remove("./log/checkimage.log")
	f, err := os.OpenFile("./log/checkimage.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger := log.New(f, "prefix", log.LstdFlags)

	dupcount = 0
	for i = 0; i < dupsize; i++ {
		for j := i + 1; j < dupsize; j++ {
			if dupArr[i] == dupArr[j] {
				logger.Println(i, "Duplicates", dupArr[i])
				dupcount++
				break
			}
		}
	}
	return dupcount
}

func check() {
	dupcount = countDuplicates(media.Media, len(media.Media))
	fmt.Println("\nThe Total Number of Duplicates in dupArr = ", dupcount)
}
