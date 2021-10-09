package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/baalamurgan/crud-golang/pkg/config"
	"github.com/baalamurgan/crud-golang/pkg/models"
	"github.com/baalamurgan/crud-golang/pkg/routes"
	"github.com/baalamurgan/crud-golang/pkg/utils"
	guuid "github.com/google/uuid"
	"github.com/gorilla/mux"
)

func Test_getUser(t *testing.T) {
	config.Connect()
	r := mux.NewRouter()
	routes.RegisterWeatherRoutes(r)

	http.Handle("/", r)

	ts := httptest.NewServer(r)
	defer ts.Close()

	newUser := models.User{
		ID:        guuid.New().String(),
		Name:      "test",
		Email:     "test@gamil.com",
		Password:  utils.HashPassword("test"),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	postBody, _ := json.Marshal(newUser)
	responseBody := bytes.NewBuffer(postBody)

	res, err := http.Post(ts.URL+"/user", "application/json", responseBody)

	if err != nil {
		t.Errorf("Expected nil, received %s", err.Error())
	}
	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected %d, received %d", http.StatusOK, res.StatusCode)
	}
}
