package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello, world!")
	say(os.Args[1])
}
func say(msg string) {
	fmt.Println("hello:" + msg)
}
