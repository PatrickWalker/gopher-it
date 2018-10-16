package main

import "flag"
import "fmt"

func main() {

	name := flag.String("name", "Ric Flair", "what yo name")

	flag.Parse()

	fmt.Printf("Hello %s! \n", *name)
}
