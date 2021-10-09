package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/baalamurgan/crud-golang/pkg/models"
	"github.com/baalamurgan/crud-golang/pkg/utils"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var Post_Collection *mongo.Collection

func PostCollection(c *mongo.Database) {
	Post_Collection = c.Collection("posts")
}
func CreatePost(w http.ResponseWriter, r *http.Request) {

	var post models.Post

	err := json.NewDecoder(r.Body).Decode(&post) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	id := guuid.New().String()

	newPost := models.Post{
		ID:        id,
		UserID:    post.UserID,
		Caption:   post.Caption,
		ImgUrl:    post.ImgUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertResult, err := Post_Collection.InsertOne(context.TODO(), newPost)

	if err != nil {
		fmt.Println(err)
	}

	res, _ := json.Marshal(insertResult)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetPostById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	postId := vars["postID"]

	post := models.Post{}

	err := Post_Collection.FindOne(context.TODO(), bson.M{"id": postId}).Decode(&post)

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		json.NewEncoder(w).Encode(utils.Message{"404", "Post not found"})
		return
	}

	res, _ := json.Marshal(post)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetPostByUserId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userID := vars["userID"]

	posts := []models.Post{}

	cursor, err := Post_Collection.Find(context.TODO(), bson.M{"userid": userID})

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		json.NewEncoder(w).Encode(utils.Message{"404", "Post not found"})
		return
	}

	for cursor.Next(context.TODO()) {
		var post models.Post
		cursor.Decode(&post)
		posts = append(posts, post) // this is like push()  oushing to an attay? yes
	}

	res, _ := json.Marshal(posts)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
