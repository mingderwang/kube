package main

import (
	"fmt"
	"log"
	"os/exec"
)

// https://gist.githubusercontent.com/anonymous/04fd2287b1912e26ec46/raw/my-redis-rc.yaml

func getFile(tag string) {
	var kube = Kube{}
	getFileRequest(tag, &kube)
	urlGist := "https://gist.githubusercontent.com/anonymous/" + kube.Gistid + "/raw/" + kube.Filename
	command1 := "wget"
	out, err := exec.Command(command1, urlGist).Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The file is download. %s\n", out)
}
