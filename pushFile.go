package main

import (
	"encoding/json"
	"github.com/davecgh/go-spew/spew"
	"io/ioutil"
	"log"
)

type Gist struct {
	Files       map[string]GistFile `json:"files"`
	Public      bool                `json:"public"`
	Description string              `json:"description"`
}

type GistFile struct {
	Content string `json:"content"`
}

func pushFile(filename string, tag string, description string) {
	println("pushing file....")
	files := map[string]GistFile{}

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal("Read file error: ", err)
	}
	files[filename] = GistFile{string(content)}

	gist := Gist{
		files,
		true,
		description,
	}
	data, err := json.Marshal(gist)
	if err != nil {
		log.Fatal("JSON Error: ", err)
	}
	spew.Dump(gist)
	spew.Dump(data)
}
