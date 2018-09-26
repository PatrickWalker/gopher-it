package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Echo Bot is ready. Type anything and err i'll echo it")
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
