package main

func searchTag(keyword string) {
	url := "http://api.log4security.com:32569/search/kube/" + keyword
	getRequestResponse(url)
}
