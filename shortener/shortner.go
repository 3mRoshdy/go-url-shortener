package shortner

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/3mRoshdy/url-shortener/generator"
)

type RequestBody struct {
	Url string
}

type ResponseBody struct {
	Url      string
	ShortUrl string
}

type URLShortner struct {
	Urls map[string]string
}

type IShortner interface {
	HandleShorten(http.ResponseWriter, *http.Request)
	HandleRedirect(http.ResponseWriter, *http.Request)
}

func (us *URLShortner) HandleShorten(w http.ResponseWriter, r *http.Request) {
	var requestBody RequestBody

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid Request method", http.StatusMethodNotAllowed)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		http.Error(w, "Failed to parse body", http.StatusBadRequest)
		return
	}

	if requestBody.Url == "" {
		http.Error(w, "URL paramater is missing", http.StatusBadRequest)
		return
	}

	shortKey := generator.GenerateShortKey()
	us.Urls[shortKey] = requestBody.Url

	fmt.Printf("Urls: %v \n", us.Urls)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(&ResponseBody{Url: requestBody.Url, ShortUrl: shortKey})
}

func (us *URLShortner) HandleRedirect(w http.ResponseWriter, r *http.Request) {
	shortKey := r.URL.Path[len("/short/"):]

	if shortKey == "" {
		http.Error(w, "Shortened key is missing", http.StatusBadRequest)
		return
	}

	// Retrieve the original URL from the `urls` map using the shortened key
	originalURL, found := us.Urls[shortKey]
	if !found {
		http.Error(w, "Shortened key not found", http.StatusNotFound)
		return
	}

	fmt.Printf("Redirecting to: %s \n", originalURL)
	http.Redirect(w, r, originalURL, http.StatusMovedPermanently)
}
