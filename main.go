package main

import (
	"fmt"
	"log"
	"net/http"
)

// using  newmux
// create a server function to handle request and write response
func HelloServer(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "hello World")
}

func main() {
	mux := http.NewServeMux()

	//handle the server function with mux router

	mux.HandleFunc("/", HelloServer)

	log.Println("listening on Port :8080")

	// address port to handle above server
	http.ListenAndServe(":8080", mux)

}
