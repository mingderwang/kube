package main

func searchTag(keyword string) {
	url := "http://api.log4security.com:8080/search/kube/" + keyword
	getRequestResponse(url)
}
