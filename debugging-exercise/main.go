package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	time.Sleep(5 * time.Second)
	setup()
	http.HandleFunc("/", hello)
	fmt.Println("Starting server, listening...")
	err := http.ListenAndServe(":"+os.Getenv("PORT"), nil)
	if err != nil {
		panic(err)
	}
}

func hello(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello, app is running ~")
}
