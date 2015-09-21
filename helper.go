package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"log"
	"net/http"
	"strconv"
	"unicode/utf8"
)

//return "" when fail, return gist id when upload success
func pushGist(gist Gist) string {
	const gistURL string = "https://api.github.com/gists"
	result := pushRequest(gist, gistURL)
	return result
}

func Short(s string, i int) string {
	if len(s) < i {
		return s
	}
	if utf8.ValidString(s[:i]) {
		return s[:i]
	}
	return s[:i-1]
}

func getRequestResponse(url string) {
	_, err1 := http.Get(url)
	if err1 != nil {
		log.Fatal("post error: ", err1)
	}
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("post error: ", err)
	}
	defer response.Body.Close()
	//spew.Dump(response.Body)
	dec := json.NewDecoder(response.Body)

	// read open bracket
	_, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}

	var buffer bytes.Buffer

	for i := 0; i < 85; i++ {
		if i == 0 {
			buffer.WriteString("NAME")
		} else if i == 27 {
			buffer.WriteString("DESCRIPTION")
		} else if i == 80 {
			buffer.WriteString("STARS")
		} else {
			buffer.WriteString(" ")
		}
	}
	fmt.Println(buffer.String())
	var kube Kube
	// while the array contains values
	for dec.More() {

		// decode an array value (Message)
		err := dec.Decode(&kube)
		if err != nil {
			log.Fatal(err)
		}

		var buffer bytes.Buffer

		for i := 0; i < 85; i++ {
			if i == 0 {
				s1 := Short(kube.Tag, 25)
				i = i + len(s1) - 4
				buffer.WriteString(s1)
			} else if i == 27 {
				s2 := Short(kube.Description, 51)
				i = i + len(s2) - 11
				buffer.WriteString(s2)
			} else if i == 80 {
				buffer.WriteString(strconv.Itoa(kube.Like))
			} else {
				buffer.WriteString(" ")
			}
		}
		fmt.Println(buffer.String())
	}
}

func getRequest(url string) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("post error: ", err)
	}
	defer response.Body.Close()
}

func pushRequest(gist interface{}, url string) string {
	var obj map[string]interface{}

	jsonData, err := json.Marshal(gist)
	if err != nil {
		log.Fatal("push gist marshel error: ", err)
	}

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
