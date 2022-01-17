package main

import (
	"fmt"
	"net/http"
)

// create a server function to handle request and write resonse
func HelloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello World")
}

func main() {

	http.HandleFunc("/", HelloServer)

	// address port to handle above server
	http.ListenAndServe(":8080", nil)

}
