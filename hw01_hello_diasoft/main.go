package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

func main() {
	text := "Hello, DIASOFT!"

	reverseText := reverse.String(text)

	fmt.Println(reverseText)
}
