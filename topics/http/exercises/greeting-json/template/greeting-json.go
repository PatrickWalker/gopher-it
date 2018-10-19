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

//NameRequest contains the name we're looking to change the greeting to use
type NameRequest struct {
	Name string
}

//NameResponse contains the status of the update
type NameResponse struct {
	Status string `json:"status"`
}

//Greeting contains a greeting Message. Could also include Language
type Greeting struct {
	Message string
}

//https://golang.org/doc/articles/wiki/ talks about building web apps in golang
//sayhelloName handles 1 part of the service we're building which is the go endponit
func sayhelloName(w http.ResponseWriter, r *http.Request) {

}

//nameset is the endpoint that allows us to change who we say hello to
func nameSet(w http.ResponseWriter, r *http.Request) {

}

func main() {
	//default value so if get hit first the greeting makes sense
	name = "world"
	//this is just so we know something is running
	fmt.Println("Listening on Port 9091")
	//this is the importsnt but http is listening on port 9090
	//can set a handler here too but we don't need more
	err := http.ListenAndServe(":9091", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
