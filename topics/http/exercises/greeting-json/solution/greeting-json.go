//we're making a web server which has 2 routes
// /hello which is a go which says hello {{name}}
// /name which is a http post operation which changes the name
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var name string

//NameRequest contains the name we're looking to change the greeting to use
type NameRequest struct {
	Name string `json:"name"`
}

//NameResponse contains the status of the update
type NameResponse struct {
	Status string `json:"status"`
}

//Greeting contains a greeting Message. Could also include Language
type Greeting struct {
	Message string `json:"greeting"`
}

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
	w.Header().Set("Content-Type", "application/json")
	greet := Greeting{Message: fmt.Sprintf("Hello %s!", name)}

	//this is a sneaky way to do it
	json.NewEncoder(w).Encode(greet)

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
	var nameResp NameResponse
	nameResp = NameResponse{}
	//Reads the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		//http status 400 means you've sent us something bad
		w.WriteHeader(http.StatusBadRequest)
		nameResp.Status = "Fail"
		json.NewEncoder(w).Encode(nameResp)
		return
	}

	var nameReq NameRequest
	err = json.Unmarshal(body, &nameReq)
	if err != nil {
		log.Println(err.Error())
	}
	//change the package variable
	name = nameReq.Name
	//response again is very basic
	nameResp.Status = "Ok"
	json.NewEncoder(w).Encode(nameResp)
}

func main() {
	//registers the functions to the paths
	http.HandleFunc("/hello", sayhelloName) // set router path
	http.HandleFunc("/name", nameSet)       // set router path
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
