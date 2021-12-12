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
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type Poll struct {
	ID    primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Title string             `json:"title,omitempty" bson:"title,omitempty"`
}

var client *mongo.Client

func DefaultEndpoint(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintln(response, "Works so far")
}

func CreatePollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	var poll Poll
	json.NewDecoder(request.Body).Decode(&poll)
	collection := client.Database("myDB").Collection("polls")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	result, err := collection.InsertOne(ctx, poll)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
	fmt.Println(result)
}

func GetPollEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Add("content-type", "application/json")
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	var poll Poll
	collection := client.Database("myDB").Collection("polls")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err := collection.FindOne(ctx, Poll{ID: id}).Decode(&poll)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(poll)
	fmt.Println(poll, poll.ID, poll.Title)
}

// https://stackoverflow.com/questions/54778520/mongo-go-driver-failing-to-connect
func ConnectMongo() {
	var (
		client     *mongo.Client
		mongoURL = "mongodb://localhost:27017"
	)
 
	// Initialize a new mongo client with options
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURL))
 
	// Connect the mongo client to the MongoDB server
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
 
	// Ping MongoDB
	ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
	if err = client.Ping(ctx, readpref.Primary()); err != nil {
		fmt.Println("could not ping to mongo db service: %v\n", err)
		return
	}
 
	fmt.Println("Connected to nosql database:", mongoURL)
 
 }

func main() {
	ConnectMongo()

	router := mux.NewRouter()
	const port string = ":8080"
	router.HandleFunc("/", DefaultEndpoint).Methods("GET")
	router.HandleFunc("/poll", CreatePollEndpoint).Methods("POST")
	router.HandleFunc("/poll/{id}", CreatePollEndpoint).Methods("GET")

	log.Print("Server listening on port ", port)
	log.Fatalln(http.ListenAndServe(port, router))
}
