package main

import (
	"log"
	"net/http"

	cont "gorest/controller"

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
	router.HandleFunc("/post", cont.GetPosts).Methods(("GET"))
	router.HandleFunc("/post", cont.AddPost).Methods(("POST"))
	router.HandleFunc("/post/{postId}", cont.UpdatePost).Methods(("PUT"))
	router.HandleFunc("/post/{postId}", cont.DeletePost).Methods(("DELETE"))

	log.Println("Listening on port", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
