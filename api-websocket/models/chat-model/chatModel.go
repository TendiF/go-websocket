package chatModel

import (
	"time"
	"os"
	"context"
	"log"
	"encoding/json"
	. "go-chat/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/bson"
	. "github.com/gobeam/mongo-go-pagination"
)

func Add(chat Chat) Chat{
	chat.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection("chats")
	if insertResult, err := collection.InsertOne(context.TODO(), chat); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			chat.ID = oid
		}
	} else {
		log.Fatal(err)
	}
	return chat
}

func Get(page int64, limit int64) []byte{
	filter := bson.M{}
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
	return b
}	