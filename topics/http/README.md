
# Still work in progress. Liable to change especially the solutions


# HTTP 
![enter image description here](https://media1.tenor.com/images/d1aa4ef7a0a3b9d8b44a3f8b4613238c/tenor.gif)

## HTTP Package

Like a lot of the things we've spoken about Go ships a http package as part of the standard lib. You can use some of the other packages/frameworks such as [echo](https://github.com/labstack/echo) and others which adds another layer on top to make it even easier.

Today we're going to use the standard lib but don't worry what you learn here is pretty consistent across most frameworks 

### Interfaces

Handler is the most important one
```go
type Handler interface {
    ServeHTTP(ResponseWriter, *Request)
}
```
It's responsible for understanding and parsing info from the request object and writing it to the ResponseWriter which itself is an interface

```go
type ResponseWriter interface {
    Header() Header
    Write([]byte) (int, error)
    WriteHeader(int)
}
```



## Http Client

## Http Server

### Http Handlers/ Middleware
Thats because stdin, stderr and stdout are open files
```
package main

  

import (

"bufio"

"fmt"

"os"

)

  

func  main() {

scanner  := bufio.NewScanner(os.Stdin)

fmt.Println("Echo Bot is ready. Type anything and err i'll echo it")

for scanner.Scan() {

fmt.Println(scanner.Text())

}

}
```
The above example uses a buf scanner to read from Stdin. os.StdIn is essentially a file so implements io.Reader


# Exercise 1 
Build a web service that greets people.

There should be two routes 
GET /hello - returns a string that says hello <name>
POST /name - updates the name returned by the get endpoint

|  | Source | PlayGround| 
|--|--| -- | 
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |

# JSON 

## JSON Processing
### Marshalling

### UnMarshalling

### Tags

# Exercise 2
Update your endpoint to take a JSON object rather than a plain text and return one too

|  | Source | PlayGround| 
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

|  | Source | PlayGround| 
|--|--| -- | 
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |