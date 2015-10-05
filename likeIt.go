package main

func likeIt(tag string) {
	url := "http://api.log4security.com:31819/like/kube/" + tag
	getRequest(url)
}
