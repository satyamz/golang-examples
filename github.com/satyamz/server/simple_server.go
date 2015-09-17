package main

import (
    "fmt"
    "net/http"                                             //package for http based web programs
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Println("We are in handler")
    fmt.Fprintf(w, "Welcome to the webserver!")
}

func main() {
    http.HandleFunc("/", handler)                          // redirect all urls to the handler function
    http.ListenAndServe(":9999", nil)                      // listen for connections at port 9999 on the local machine
}
