package main

import (
	"fmt"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func main() {
	port := os.Getenv("PORT")
	fmt.Println("Obtained port from Env $PORT:", port)

	fmt.Fprintf(os.Stderr, "App requires minimum 128M to run, only 16M is allocated currently")

	http.HandleFunc("/", handler)
	http.ListenAndServe(":"+port, nil)
}
