package main

func likeIt(tag string) {
	url := "http://api.log4security.com:8080/like/kube/" + tag
	getRequest(url)
}
