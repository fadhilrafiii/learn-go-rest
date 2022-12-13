package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const (
	port = ":9000"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte(`Hello`))
	})
	router.HandleFunc("/post", GetPost)
	log.Println("Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
