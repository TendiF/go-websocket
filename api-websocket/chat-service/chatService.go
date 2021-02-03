package chatService

import (
	"os"
	"context"
	"encoding/json"
	"net/http"
	"log"
	. "go-chat/utils"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method)

	if r.Method == "GET" {
		getChat(w, r)
		return
	}

	if r.Method == "POST" {
		AddChat(w, r)
		return
	}
}

func getChat(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hello"))
}

func AddChat(w http.ResponseWriter, r *http.Request){
	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
	}

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection("chats")
	insertResult, err := collection.InsertOne(context.TODO(), chat)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted post with ID:", insertResult.InsertedID)
	w.Write([]byte("POST"))
}
