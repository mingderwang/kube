package main

func searchTag(keyword string) {
	url := "http://api.log4security.com:31819/search/kube/" + keyword
	getRequestResponse(url)
}
