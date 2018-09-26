
  

  

# Handling Input

  
![enter image description here](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSaco9Rw_0PnUQqyW_AnREYc4U_21Yr7hWXk0msnXDwaaP8Mms2)
  

## Command Line Tools
Ok we're going to get a little more involved than we have so far. You now know more than enough that you can build some righteous Go stuff but most systems need some input to do interesting things

### Flags

Parsing input values for your command line project can be a bit awkward. Quite a lot of boiler plater code etc so it's nice that Go ship something with the [standard lib](https://golang.org/pkg/flag/) to help with that

```go
import "flag" 
func main() {

//Basic flag declarations are available for string, integer, and boolean options. Here we declare a string flag  `word`  with a default value  `"foo"`  and a short description. This  `flag.String`  function returns a string pointer (not a string value); we’ll see how to use this pointer below.

    wordPtr := flag.String("word", "foo", "a string")

//This declares  `numb`  and  `fork`  flags, using a similar approach to the  `word`  flag.

    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("fork", false, "a bool")

//It’s also possible to declare an option that uses an existing var declared elsewhere in the program. Note that we need to pass in a pointer to the flag declaration function.

    var svar string
    flag.StringVar(&svar, "svar", "bar", "a string var")

//Once all flags are declared, call  `flag.Parse()`  to execute the command-line parsing.

    flag.Parse()
}
```
So we can enforce types at the command line and set defaults etc as well. 


### Reading files
Probably more common than reading keyboard input is reading files especially for things like batch processing etc.

There's a few ways to do it but before that we'll talk about something that is really important with pretty much all reading and that's io.Reader

```go
type Reader interface {  
//populates the byte array
//n is number of bytes read
//err is if something went wrong
//io.EOF err is the standard way to represent the end of the data
  Read(p []byte) (n int, err error)  
}
```
io.Reader is everywhere. It's used for files, handy in this section ;), stdin, http request bodies. Having such a clean and small interface means all those pretty disparate things above can be handled in the same way.

[this blog](https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b) is a love letter to io.reader but it also talks about some pretty class use cases like using a gzip reader to wrap a http request body or a file reader to decompress as you read. 
[this one](https://medium.com/learning-the-go-programming-language/streaming-io-in-go-d93507931185) covers more detail on io.readers and writers

Let's look at a few options with files. Examples taken from [Go by Example](https://gobyexample.com/reading-files)

```go
  

[![](https://gobyexample.com/play.png "Run code")](http://play.golang.org/p/2kEKXq-kUV)

package main

import (
    "io"
    "io/ioutil"
)
func main() {
    dat, err := ioutil.ReadFile("/tmp/dat")
    check(err)
    //...much more stuff
}
```
So this works and is great but there is a problem with it. We're reading the whole file contents into memory which is ok if the file is small but what if this is the transaction list for a store and its the week before Christmas... 

bufio to the rescue. bufio offers either a buffered reader or scanner

```go
package main 
import  (  "io"  "io/ioutil"  ) 
 func  main()  {
    f, err := os.Open("/tmp/dat")
    check(err)
    //os.File is used as the input to create a reader
    r4 := bufio.NewReader(f)
    //peek fetches some bytes without advancing the reader
    b4, err := r4.Peek(5)
    //you can also read an individual or set of bytes or lines
//scanners allow you to parse on things like lines,words or runes
	
    scanner := bufio.NewScanner(f)
    //wow this is easy it is based on tokens found by the split function
    //split defaults to new lines but you can provide your own impl
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
```
Also Go has inbuilt support for csv, json and xml parsing in the [encoding package](https://golang.org/pkg/encoding/)

### Reading keyboard input

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
Theres a number csv [files](exercises/making-the-grade/files) in the git repo

It has a list of students with name, student number and their mark for that subject.

Write a go program which takes a command line argument (flag) the subject name to load and then reads in that file. Then print out the student number and their Grade using the following grade sheet


|Score|Grade  |
|--|--|
|90-100|A+|
| 80-89 | A |
|70-79  | B |
| 60-69 | C |
| 50-59 | D |
| 0-49 | F |

The files are valid but do think about what you might to do protect your code from issues from badly formatted files. There's also no header on the files.

|  | Source | PlayGround| 
|--|--| -- | 
| Template | [source](exercises/making-the-grade/template/grade.go) | [pg](https://play.golang.org/p/xrCU0H-oxYC) |
| Solution |[source](exercises/making-the-grade/solution/grade.go) | [pg](https://play.golang.org/p/maeSBOLEyNJ) |
