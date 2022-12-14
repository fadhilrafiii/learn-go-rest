package repository

import (
	"context"
	"errors"
	"fmt"
	e "gorest/entity"
	"log"
)

type PostRepository interface {
	Save(post *e.Post) (*e.Post, error)
	FindAll() ([]e.Post, error)
	Update(postId string, payload map[string]interface{}) (*e.Post, error)
	Delete(postId string) error
}

type repo struct{}

func NewRepository() PostRepository {
	return &repo{}
}

func (*repo) Save(post *e.Post) (*e.Post, error) {
	ctx := context.Background()
	store, _ := InitFirestore(ctx)

	defer store.Close()
	fmt.Println(post)
	data, _, err := store.Collection("posts").Add(ctx, map[string]interface{}{
		"Title": post.Title,
	})

	if err != nil {
		log.Fatalln("Error saving post data to firestore")
		return nil, err
	}

	result := e.Post{
		Id:    data.ID,
		Title: post.Title,
	}

	return &result, nil
}

func (*repo) FindAll() ([]e.Post, error) {
	ctx := context.Background()
	store, _ := InitFirestore(ctx)
	var posts []e.Post = []e.Post{}

	defer store.Close()
	collections, err := store.Collection("posts").Documents(ctx).GetAll()

	if err != nil {
		log.Fatalln("Error getting posts data")
		return nil, err
	}

	for _, v := range collections {
		doc, _ := v.Ref.Get(ctx)
		docData := doc.Data()

		post := e.Post{
			Id:    doc.Ref.ID,
			Title: docData["Title"].(string),
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func (*repo) Update(postId string, payload map[string]interface{}) (*e.Post, error) {
	ctx := context.Background()
	store, err := InitFirestore(ctx)

	if err != nil {
		log.Fatalln("Failed to initialize Firestore!")
		return nil, err
	}

	stringPath := fmt.Sprintf("posts/%s", postId)

	doc := store.Doc(stringPath)
	_, errData := doc.Get(ctx)

	if errData != nil {
		return nil, errors.New("Post data not found!")
	}

	_, err2 := doc.Set(ctx, map[string]interface{}{
		"Title": payload["title"],
	})

	if err2 != nil {
		log.Fatalln("Failed to update post data!")
		return nil, err2
	}

	return &e.Post{
		Id:    doc.ID,
		Title: payload["title"].(string),
	}, err
}

func (*repo) Delete(postId string) error {
	ctx := context.Background()
	store, err := InitFirestore(ctx)

	if err != nil {
		log.Fatalln("Failed to initialize Firestore!")

		return err
	}

	stringPath := fmt.Sprintf("posts/%s", postId)
	doc := store.Doc(stringPath)
	_, errData := doc.Get(ctx)

	fmt.Print("errData", errData)

	if errData != nil {
		return nil
	}

	_, errDelete := doc.Delete(ctx)

	fmt.Print("errDelete", errDelete)

	if errDelete != nil {
		return errDelete
	}

	return nil
}
