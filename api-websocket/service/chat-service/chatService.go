package chatService

import (
	"os"
	"time"
	"strconv"
	"context"
	"encoding/json"
	"net/http"
	"log"
	. "go-chat/utils"
	. "github.com/gobeam/mongo-go-pagination"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	// Example for Normal Find query
	filter := bson.M{}
	var limit int64 = 10
	var page int64 = 1

	if keys, ok := r.URL.Query()["page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		page = i
	}

	if keys, ok := r.URL.Query()["per_page"]; ok {
		i , err :=  strconv.ParseInt(keys[0], 10, 64)
		if err != nil {
			log.Println(err)
		}
		limit = i
	}

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection("chats")
	projection := bson.D{
	}
	
	paginatedData, err := New(collection).Limit(limit).Page(page).Sort("price", -1).Select(projection).Filter(filter).Find()
	if err != nil {
		panic(err)
	}

	var chats []Chat
	for _, raw := range paginatedData.Data {
		var chat *Chat
		if marshallErr := bson.Unmarshal(raw, &chat); marshallErr == nil {
			chats = append(chats, *chat)
		}

	}

	b, err := json.Marshal(map[string]interface{}{
		"data" : chats,
		"pagination" : paginatedData.Pagination,
	})

	if err != nil {
		log.Println("error:", err)
	}

	w.Write([]byte(b))
}

func AddChat(w http.ResponseWriter, r *http.Request){
	var chat Chat
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil {
		log.Println(err)
	}

	chat.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection("chats")
	if insertResult, err := collection.InsertOne(context.TODO(), chat); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			chat.ID = oid
		}
	} else {
		log.Fatal(err)
	}

	b, err := json.Marshal(chat)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}
