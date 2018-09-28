//we're making a web server which has 2 routes
// /hello which is a go which says hello {{name}}
// /name which is a http post operation which changes the name
package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
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

//listShort is the endpoint that allows us list all the shortened URLs
func listShort(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only GET supported at this endpoint"))
		return
	}
	json.NewEncoder(w).Encode(shortenedURL)
	return
}

//shortenURL is the endpoint that allows us to change who we say hello to
func shortenURL(w http.ResponseWriter, r *http.Request) {
	//if not a post we're not proceeding
	//we could handle a GET differently from a POST etc if we wanted to
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST supported at this endpoint"))
		return
	}

	//Reads the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//http status 400 means you've sent us something bad
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request"))
		return
	}
	var shortReq ShortenRequest
	err = json.Unmarshal(body, &shortReq)
	if err != nil {
		log.Println(err.Error())
	}

	//create the response object
	var shortResp ShortenResponse
	shortResp = ShortenResponse{OriginalURL: shortReq.URL}

	//check if it exists if so fetch the existing value and return
	exists, err := shortenedExist(shortReq.URL)
	if exists {
		short, err := getShortenedURL(shortReq.URL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Unable to shortern url"))
			return
		}
		shortResp.NewURL = short
		json.NewEncoder(w).Encode(shortResp)
		return
	}
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to shortern url"))
		return
	}

	//if not generate the new shortened URL
	short, err := getShortToken()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to shortern url"))
		return
	}

	err = storeShortened(short, shortReq.URL)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Unable to shortern url"))
		return
	}
	shortResp.NewURL = short
	json.NewEncoder(w).Encode(shortResp)
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

//storeShortened places the token and url pair in the map
//this could also be done to a file or db
func storeShortened(shortened string, original string) error {
	shortenedURL[shortened] = original
	urlShortenLookup[original] = shortened
	return nil
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

//getShortenedURL takes a full URL and if its in the map knows its been shortened
func getShortenedURL(url string) (string, error) {
	//this map is keyed on the originl urls so we dont have to traverse the other
	if val, ok := urlShortenLookup[url]; ok {
		return val, nil
	}
	return "", errors.New("URL doesn't exist")
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
	//registers the functions to the paths
	http.HandleFunc("/shorten", shortenURL) // set router path
	http.HandleFunc("/list", listShort)     // set router path
	http.HandleFunc("/", redirect)          // set router path

	//initializes the maps
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
