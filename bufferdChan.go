package main

import "fmt"

func main(){
	messages := make(chan string, 2)
	messages <- "buffed"
	messages <- "1"

	// messages <- "1"

	fmt.Println(<-messages)
}