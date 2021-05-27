package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Media struct {
	Media []File `json:"data"`
}

type File struct {
	Id   string `json:"id"`
	Path string `json:"path"`
}

var media Media

func loadjson(path string) bool {
	// Open our jsonFile
	jsonFile, err := os.Open(path)

	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)

		return false
	}
	fmt.Println("Successfully Opened", path)

	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &media)

	return true
}
