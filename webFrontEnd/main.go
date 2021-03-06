package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	initClient()
}

func main() {
	host := port()
	fmt.Println("Serving at port ", host)

	log.Fatal(http.ListenAndServe(host, handler()))
}

func handler() http.Handler {
	r := http.NewServeMux()
	tem := templateHandler{fileName: "tweet.html"}
	r.Handle("/index", &tem)
	r.HandleFunc("/tweet", protectTweet(serveTweet))
	r.HandleFunc("/retweets", protectTweets(serveTopTweets))

	return r
}

func port() string {
	return ":8080"
}
