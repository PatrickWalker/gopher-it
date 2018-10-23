
  

 
# HTTP

![enter image description here](https://media1.tenor.com/images/d1aa4ef7a0a3b9d8b44a3f8b4613238c/tenor.gif)

  

## HTTP Package

  

Like a lot of the things we've spoken about Go ships a http package as part of the standard lib. You can use some of the other packages/frameworks such as [echo](https://github.com/labstack/echo) and others which adds another layer on top to make it even easier.

  

Today we're going to use the standard lib but don't worry what you learn here is pretty consistent across most frameworks

  

### Interfaces

Interfaces play a pretty important role in http handling in Go. It's this common set of interfaces which allow for frameworks to operate in pretty consistent ways  

Handler is the most important one

```go

type  Handler  interface {

ServeHTTP(ResponseWriter, *Request)

}

```

It's responsible for understanding and parsing info from the request object and writing it to the ResponseWriter which itself is an interface

  

```go

type  ResponseWriter  interface {

Header() Header

Write([]byte) (int, error)

WriteHeader(int)

}

```

  
  
  

## Http Client
 An Http Client is something that makes HTTP calls to another destination. Usage for this might be if your service has a dependency on other api's or similar.
 
 An example from [go by example](https://dlintw.github.io/gobyexample/public/http-client.html) below
```go
// Go contains rich function for grab web contents. _net/http_ is the major
// library.
// Ref: [golang.org](http://golang.org/pkg/net/http/#pkg-examples).
package main

import (
	"net/http"
	"net/url"
)
import "io/ioutil"
import "fmt"
import "strings"

// keep first n lines
func keepLines(s string, n int) string {
	result := strings.Join(strings.Split(s, "\n")[:n], "\n")
	return strings.Replace(result, "\r", "", -1)
}

func main() {
	// We can use GET form to get result.
	//Robots is used for crawlers like search enginer indexers.
	//try the url in your browser as the Get operation is a read
	resp, err := http.Get("http://g.cn/robots.txt")
	if err != nil {
		panic(err)
	}
	//defer the closure of the response body
	defer resp.Body.Close()
	//our friend ioutil below (from handling input) reads the body
	body, err := ioutil.ReadAll(resp.Body)
	//gives us a byte array so we'll cast to string
	fmt.Println("get:\n", keepLines(string(body), 3))

	// We can use POST form to get result, too.
	//Post is sending data to an endpoint
	//Below this will be us sending a search or query request for github
	
	resp, err = http.PostForm("http://duckduckgo.com",
		url.Values{"q": {"github"}})
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err = ioutil.ReadAll(resp.Body)
	fmt.Println("post:\n", keepLines(string(body), 3))
}

```
At a high level that covers http client. Feel free to check out the std lib docs [here](https://golang.org/pkg/net/http/)

It's important to note that the above playground link won't work because the playground limits access to net/http so people don't use it as a massive crawler/scraper network. Feel free to copy the file out to a local folder and build and run.

*Bonus tip* - I find this interesting anyway if you add ".json" to reddit urls you can turn them into api calls https://www.reddit.com/r/soccerstreams.json returns the json equivalent of going direct to https://www.reddit.com/r/soccerstreams so if you don't want to get a robots.txt go have a look at what you can get out of Reddit

## Http Server

HTTP Server is a process that listens for http requests and then returns http responses for them. 

### Http Handlers/ Middleware

Quick example shown below

```
package main

import (
	"fmt"
	//the std lib package responsible for all http goodness
	"net/http"
)

func main() {
	//handlefunc taks a path "/" in this case and a function and registers that function to be called for any invocation of the path
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, you've requested: %s\n", r.URL.Path)
	})

	//this is what makes things active. This listens on port 80
	//the nil is another handler that can be passed if nil the default is used
	//default will use the HandleFunc we have already registered in the global http object
	http.ListenAndServe(":80", nil)
}

```

The above example registers an anonymous function to handle the route "/" in this case it would be localhost/ and prints out some text to explain the path that's been asked for.

You can also use named functions as well. The only thing that has to be for HandleFunc is the function should have that signature above. A handler can be used as in place of this function but the method is a bit different it's signature is shown below

```go
func  Handle(pattern string, handler Handler) { DefaultServeMux.Handle(pattern, handler) }

  

// HandleFunc registers the handler function for the given pattern

// in the DefaultServeMux.

// The documentation for ServeMux explains how patterns are matched.

func  HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {

DefaultServeMux.HandleFunc(pattern, handler)

}
```

Example below

```go
http.Handle("/foo", fooHandler)

http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
})

log.Fatal(http.ListenAndServe(":8080", nil))
```

# Testing your exercises
I recommend [postman](https://www.getpostman.com/) for this but if you're feeling brave you could write the http client as shown above in notes.

The postman collection I use is in the repo [here](exercises/gopher-it.postman_collection.json)

This is a pretty basic postman collection. It's possible to put tests and assertions in it and it's nice. You can run the individual commands. If you have problems getting postman going give me a shout :) 

# Exercise 1

Build a web service that greets people.

 
There should be two routes

GET /hello - returns a string that says hello <name>

POST /name - updates the name returned by the get endpoint

  

| | Source | PlayGround|
|--|--| -- |
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |

**By default this is running on port 9090 so if you see funny results there's a chance something else may be listening on that port (like prometheus) so if you see weirdness change the value from 9090 to something else**

# JSON

 We learned a bit about json tags and marshalling in week 3 in the handling input topic. 

## JSON Processing

To jog your memory the following should help
[Playground](https://play.golang.org/p/c-vPf4jyFRy)

Source code is available for that too [source]
  

# Exercise 2

Update your endpoint to take a JSON object rather than a plain text and return one too

  

| Type | Source | PlayGround|
|--|--| -- |
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |

  

# Exercise 3

  

The final part of the course!

Update your service to make a URL shortener service. The service should do the following

  

Post - Creates a new shortened link of the form localhost:8080/l/GUID

Get - List all shortened links and their original url at localhost:8080/list

Get - On that shortened link will redirect to the original URL

  

You don't have to worry about persistence right now but you can if you want.

  

| | Source | PlayGround|
|--|--| -- |
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |