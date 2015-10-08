package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//API :
type API struct {
	Message string
}

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		message := API{"Hello World"}
		output, err := json.Marshal(message)
		if err != nil {
			log.Println("Something went wrong!!")
		}
		fmt.Fprintf(w, string(output)) //Converted to string and sending response
	})
	log.Println("Listening on localhost:2000")
	http.ListenAndServe(":2000", nil)

}
