//we're making a web server which has 2 routes
// /hello which is a go which says hello {{name}}
// /name which is a http post operation which changes the name
package main

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
)

//key is the shortened URL and value is original for redirect
var shortenedURL map[string]string

//key is the original url and value is shortened
var urlShortenLookup map[string]string

//ShortenRequest contains the url we're trying to shorten
type ShortenRequest struct {
	URL string `json:"url"`
}

//ShortenResponse contains the shortened URL
type ShortenResponse struct {
	NewURL      string `json:"shortened"`
	OriginalURL string `json:"original"`
}

//Returns a random 16 character URL encoding safe string
func getShortToken() (short string, err error) {
	//creates a byte array
	uuid := make([]byte, 16)
	//fills it
	_, err = io.ReadFull(rand.Reader, uuid)
	if err != nil {
		//error unlikely but if so return
		return "", err
	}
	//we make sure the byte array is URL safe (no special characters)
	return base64.URLEncoding.EncodeToString(uuid), nil
}

//lookupShortURL takes a shortened URL and returns the original URL if exists
//error is returned if it doesn't exist
func lookupShortURL(url string) (string, error) {
	if val, ok := shortenedURL[url]; ok {
		return val, nil
	}
	return "", errors.New("Shortened URL doesn't exist")
}

//shortenedExist takes a full URL and if its in the map knows its been shortened
//leaving interface open to return an error incase this was reading a file or db
func shortenedExist(url string) (bool, error) {
	_, ok := urlShortenLookup[url]
	return ok, nil
}

//redirect is the endpoint that allows us to redirect to a shortened url if you've got a valud one
func redirect(w http.ResponseWriter, r *http.Request) {
	//check if in the map
	var short string
	path := r.URL.EscapedPath()
	if len(path) > 0 {
		short = path[1:]
	}
	//get the real url if exists err if not
	url, err := lookupShortURL(short)
	if err != nil {
		//404 if the supplied value is wrong
		w.WriteHeader(http.StatusNotFound)
		return
	}
	//this redirects to the URL
	http.Redirect(w, r, url, 301)
	return
}

func main() {

	//initializes the maps which store the urls
	shortenedURL = make(map[string]string)
	urlShortenLookup = make(map[string]string)

	//this is just so we know something is running
	fmt.Println("Listening on Port 9092")
	//this is the importsnt but http is listening on port 9090
	//can set a handler here too but we don't need more
	err := http.ListenAndServe(":9092", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
