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

var collection *mongo.Collection

func UserCollection(c *mongo.Database) {
	collection = c.Collection("users")
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json") // for adding Content-type

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user) // storing in person variable of type user
	if err != nil {
		fmt.Print(err)
	}

	id := guuid.New().String()

	newUser := models.User{
		ID:        id,
		Name:      user.Name,
		Email:     user.Email,
		Password:  utils.HashPassword(user.Password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	insertResult, err := collection.InsertOne(context.TODO(), newUser)

	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(insertResult) // return the mongodb ID of generated document

}

func GetUserbyId(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	userId := vars["userId"]

	user := models.User{}

	err := collection.FindOne(context.TODO(), bson.M{"id": userId}).Decode(&user)

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		json.NewEncoder(w).Encode(utils.Message{"404", "User not found"})
		return
	}

	res, _ := json.Marshal(user)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}
