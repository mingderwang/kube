package main

import (
	"bytes"
	"encoding/json"
	//	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
)

//return "" when fail, return gist id when upload success
func pushGist(gist Gist) string {
	const gistURL string = "https://api.github.com/gists"
	result := pushRequest(gist, gistURL)
	return result
}

func pushRequest(gist interface{}, url string) string {
	var obj map[string]interface{}

	jsonData, err := json.Marshal(gist)
	if err != nil {
		log.Fatal("push gist marshel error: ", err)
	}
	//	spew.Dump(jsonData)

	payload := bytes.NewBuffer(jsonData)
	respond, err := http.Post(url, "application/json", payload)
	if err != nil {
		log.Fatal("post error: ", err)
	}
	err = json.NewDecoder(respond.Body).Decode(&obj)
	if err != nil {
		log.Fatal("Response json error: ", err)
	}
	if str, ok := obj["id"].(string); ok {
		return str
	} else {
		log.Fatal("Response id is not a string error: ", ok)
		return ""
	}
}

func pushKube(kube Kube) string {
	const kubeURL string = "http://api.log4security.com:8080/kube"
	result := pushRequest(kube, kubeURL)
	return result
}
