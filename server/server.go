package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Poll struct {
	ID			primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title 		string             `json:"title,omitempty" bson:"title,omitempty"`
}
var client *mongo.Client

func DefaultEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Works so far")
}

func CreatePollEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Println("1")
	response.Header().Add("content-type", "application/json")
	fmt.Println("2")
	var poll Poll
	json.NewDecoder(request.Body).Decode(&poll)
	fmt.Println("3")
	fmt.Println(poll.ID, poll.Title)
	collection := client.Database("myDB").Collection("polls")
	fmt.Println("4")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	fmt.Println("5")
	result, _ := collection.InsertOne(ctx, poll)
	fmt.Println("6")
	json.NewEncoder(response).Encode(result)
}

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, "mongodb://localhost:7017")

	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", DefaultEndpoint).Methods("GET")
	router.HandleFunc("/poll", CreatePollEndpoint).Methods("POST")

	log.Print("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
