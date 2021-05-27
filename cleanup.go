package main

import (
	"log"
	"os"
	"path/filepath"
)

func cleanup(dir string) error {
	files, err := filepath.Glob(filepath.Join(dir, "*"))
	if err != nil {
		return err
	}
	for _, file := range files {
		err = os.RemoveAll(file)
		if err != nil {
			return err
		}
	}

	f, err := os.Create(dir + ".gitkeep")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	return nil
}
