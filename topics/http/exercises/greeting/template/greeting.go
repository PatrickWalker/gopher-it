//we're making a web server which has 2 routes
// /hello which is a go which says hello {{name}}
// /name which is a http post operation which changes the name
package main

import (
	"fmt"
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

}

//nameset is the endpoint that allows us to change who we say hello to
func nameSet(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//registers the functions to the paths

	//default value so if get hit first the greeting makes sense
	name = "world"
	//this is just so we know something is running
	fmt.Println("Listening on Port 9090")
	//this is the importsnt but http is listening on port 9090
	//can set a handler here too but we don't need more
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
