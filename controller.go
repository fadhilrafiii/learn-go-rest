package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

func GetPost(res http.ResponseWriter, req *http.Request) {
	ctx := context.Background()
	res.Header().Set("Content-Type", "application/json")
	posts := []Post{}

	firestoreClient, _ := InitFirestore(ctx)
	fmt.Println(firestoreClient)

	defer firestoreClient.Close()
	collections, err := firestoreClient.Collection("posts").Documents(ctx).GetAll()
	if err != nil {
		res.Write([]byte(`Error initalizing firestore!`))
	}

	for _, v := range collections {
		doc, _ := v.Ref.Get(ctx)
		docData := doc.Data()

		posts = append(posts, Post{
			Id:    docData["Id"].(int64),
			Title: docData["Title"].(string),
		})
	}

	postBytes, err := json.Marshal(posts)

	if err != nil {
		res.Write([]byte(`Error marshalling data!`))
	}

	res.WriteHeader(http.StatusOK)
	res.Write(postBytes)
}
