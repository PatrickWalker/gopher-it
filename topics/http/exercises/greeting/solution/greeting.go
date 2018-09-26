//we're making a web server which has 2 routes
// /hello which is a go which says hello {{name}}
// /name which is a http post operation which changes the name
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var name string

//https://golang.org/doc/articles/wiki/ talks about building web apps in golang
//sayhelloName handles 1 part of the service we're building which is the go endponit
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	//if the request method isn't get we throw a 405 response
	if r.Method != "GET" {
		//this means the http method/verb is not allowed at this endpoint
		w.WriteHeader(http.StatusMethodNotAllowed)
		//this is the response body
		w.Write([]byte("Only GET supported at this endpoint"))
		return

	}
	//no need to formally set the 200 status ok thats the default
	//this is just a plain text response at this stage
	fmt.Fprintf(w, "Hello %s!", name)

}

//nameset is the endpoint that allows us to change who we say hello to
func nameSet(w http.ResponseWriter, r *http.Request) {
	//if not a post we're not proceeding
	//we could handle a GET differently from a POST etc if we wanted to
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Only POST supported at this endpoint"))
		return
	}
	//normally you would also check the content-type here

	//Reads the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//http status 400 means you've sent us something bad
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Unable to parse request"))
		return
	}
	//change the package variable
	name = string(body)
	//response again is very basic
	fmt.Fprintf(w, "Name Updated")
}

func main() {
	//registers the functions to the paths
	http.HandleFunc("/hello", sayhelloName) // set router path
	http.HandleFunc("/name", nameSet)       // set router path
	//default value so if get hit first the greeting makes sense
	name = "world"
	//this is just so we know something is running
	fmt.Println("Listening on Port 9090")
	//this is the importsnt but http is listening on port 9090
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
