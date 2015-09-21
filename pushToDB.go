package main

import (
//	"github.com/davecgh/go-spew/spew"
)

type Kube struct {
	Gistid      string `json:"gistid"`
	Filename    string `json:"filename"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Like        int    `json:"like"`
}

func pushToDB(id string, filename string, tag string, description string) {
	kube := Kube{
		id,
		filename,
		description,
		tag,
		0,
	}
	pushKube(kube)
}
