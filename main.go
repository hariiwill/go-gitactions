package main

import "fmt"

var version = "dev"

func main() {
	fmt.Printf("Version : %s \n", version)
	fmt.Println(hello())

}

// Function to run hello world

func hello() string {
	return "Hello! Golang"
}
