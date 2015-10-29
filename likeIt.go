package main

func likeIt(tag string) {
	url := "http://api.log4security.com:32569/like/kube/" + tag
	getRequest(url)
}
