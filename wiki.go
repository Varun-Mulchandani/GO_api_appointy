package main

//Importing packages

import (
	"context"
	"encoding/json"
	"fmt"

	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
  "go.mongodb.org/mongo-driver/mongo/options"
)

// For Database
var client *mongo.Client

// For http routing
var mux = http.NewServeMux()

// Creating a nested meeting structure

type Meeting struct {
	ID        primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Participants Participant
	Title string             `json:"title,omitempty" bson:"title,omitempty"`
	Start_Time string `json:"start_time,omitempty" bson:"start_time,omitempty"`
	End_Time string `json:"end_time,omitempty" bson:"end_time,omitempty"`
	TS string `json:"ts,omitempty" bson:"ts,omitempty"`
}

type Participant struct {
	Name string `json:"name,omitempty" bson:"name,omitempty"`
	Email string `json:"email,omitempty" bson:"email,omitempty"`
	RSVP string `json:"rsvp,omitempty" bson:"rsvp,omitempty"`
}

// Post request to schedule the meeting
func CreateMeetingEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meeting Meeting
	_ = json.NewDecoder(request.Body).Decode(&meeting)
	collection := client.Database("appointy3").Collection("meet2")
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	result, _ := collection.InsertOne(ctx, meeting)
	json.NewEncoder(response).Encode(result)
}

//Under construction

//func GetMeetingEndpoint(response http.ResponseWriter, request *http.Request) {
//	response.Header().Set("content-type", "application/json")
//	//params := mux.Vars(request)
//	//id, _ := primitive.ObjectIDFromHex(params["id"])
//	id := request.URL.Query().Get("id")
//	id_, _ := primitive.ObjectIDFromHex(id)
//	//fmt.Println(request)
//	//fmt.Println(id)
//	var meeting Meeting
//	collection := client.Database("appointy3").Collection("meet2")
//	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
//	err := collection.FindOne(ctx, Meeting{ID: id_}).Decode(&meeting)
//	if err != nil {
//		response.WriteHeader(http.StatusInternalServerError)
//		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
//		return
//	}
//	json.NewEncoder(response).Encode(meeting)
//}

// Get request to display all the meetings
func GetMeetingsEndpoint(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var meetings []Meeting
	collection := client.Database("appointy3").Collection("meet2")
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	cursor, err := collection.Find(ctx, bson.M{})

	// Error handling
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}

	// Cursor to extract data from DB
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var meeting Meeting
		cursor.Decode(&meeting)
		meetings = append(meetings, meeting)
	}

	// Error handling
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(meetings)
}

func main() {
	//mux := http.NewServeMux()
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	// Connecting to mongoserver
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, _ = mongo.Connect(ctx, clientOptions)

	mux.HandleFunc("/meetings", CreateMeetingEndpoint)
	mux.HandleFunc("/meetingsall", GetMeetingsEndpoint)
	//mux.HandleFunc("/meetings/{id}", GetMeetingEndpoint)//.Methods("GET")
	http.ListenAndServe(":12345", mux)
}
