package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//API : define
type API struct {
	Message string "json:message"
}

//Hello :
func Hello(w http.ResponseWriter, r *http.Request) {
	urlParams := mux.Vars(r)
	log.Println("URL Parameters :", urlParams)
	name := urlParams["user"]
	HelloMessage := "Hello " + name
	output := API{HelloMessage}
	out, err := json.Marshal(output)
	if err != nil {
		log.Panic(err.Error())
	}
	w.Header().Set("pragma", "DoIt")
	fmt.Fprintf(w, string(out))
}

func main() {
	gorillaRoute := mux.NewRouter()
	gorillaRoute.HandleFunc("/api/v1/{user:[0-9]+}", Hello)
	gorillaRoute.HandleFunc("/api/", Hello)
	http.Handle("/", gorillaRoute)
	log.Println("Listening on port 4000")
	http.ListenAndServe(":4000", nil)
}
