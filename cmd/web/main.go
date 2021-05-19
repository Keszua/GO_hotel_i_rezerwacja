package main

import (
	"fmt"
	"net/http"

	"example.com/udemy/cmd/pkg/handlers"
	//"example.com/cmd/pkg/handlers"
	//"example.com/udemy/cmd/pkg/handlers"
	//"example.com/main/cmd/pkg/handlers"
)

const portNumber = ":8080"

//main is the main appliction function
func main() {
	fmt.Println("Starting aplicatiion on port", portNumber)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	_ = http.ListenAndServe(portNumber, nil)

}
