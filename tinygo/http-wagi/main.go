package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Content-Type: text/plain")
	fmt.Println("")
	fmt.Println("Hello, World Mar 7b in TinyGo world!")

	printenv()
}

func printenv() {
	//fmt.Printf("HOME=%s\n", os.Getenv("HOME"))
	//fmt.Println(os.Environ())
	for _, e := range os.Environ() {
		fmt.Printf("%s\n", e)
	}
}
