package main

import (
	"fmt"
	"net/http"

	shortner "github.com/3mRoshdy/url-shortener/shortener"
)

func main() {
	var shortener shortner.IShortner = &shortner.URLShortner{
		Urls: make(map[string]string),
	}

	http.HandleFunc("/shorten", shortener.HandleShorten)
	http.HandleFunc("/short/", shortener.HandleRedirect)

	fmt.Println("URL Shortener is running on :8080")
	http.ListenAndServe(":8080", nil)
}
