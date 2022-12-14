package controller

import (
	"encoding/json"
	"fmt"
	e "gorest/entity"
	r "gorest/repository"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func GetPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	posts, err := r.NewRepository().FindAll()

	if err != nil {
		log.Fatalln(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "error": "Error getting data!" }`))
		return
	}

	postBytes, err := json.Marshal(posts)

	if err != nil {
		log.Fatalln(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "error": "Error marshalling data!" }`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(postBytes)
}

func AddPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	var payload e.Post
	json.NewDecoder(req.Body).Decode(&payload)

	post, err := r.NewRepository().Save(&payload)

	if err != nil {
		log.Fatalln(err)
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "error": "Failed to add post!" }`))
	}

	result, err := json.Marshal(post)

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func UpdatePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(req)

	var body map[string]interface{}
	json.NewDecoder(req.Body).Decode(&body)

	fmt.Println("body", body)

	post, err := r.NewRepository().Update(vars["postId"], body)

	if post == nil {
		res.WriteHeader(http.StatusNotFound)
		res.Write(nil)
	}

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "error": "Failed to update post data" }`))
		return
	}

	result, err := json.Marshal(post)

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(`{ "error": "Failed to marshal post data" }`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func DeletePost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req)

	err := r.NewRepository().Delete(vars["postId"])

	if err != nil {
		res.WriteHeader(http.StatusNoContent)
		res.Write([]byte(`{ "message": "Data not exist already!" }`))
		return
	}

	res.WriteHeader(http.StatusOK)
	res.Write([]byte(`{ "message": "Success delete post data!" }`))
}
