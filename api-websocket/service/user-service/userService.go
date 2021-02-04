package userService

import (
	"os"
	"time"
	// "strconv"
	"context"
	"encoding/json"
	"net/http"
	"log"
	. "go-chat/utils"
	// . "github.com/gobeam/mongo-go-pagination"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Main(w http.ResponseWriter, r *http.Request){
	log.Println(r.Method)

	if r.Method == "POST" {
		AddUser(w, r)
		return
	}

}

func AddUser(w http.ResponseWriter, r *http.Request){
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Println(err)
	}

	user.CreatedAt = primitive.NewDateTimeFromTime(time.Now())

	collection := MongoClient.Database(os.Getenv("MONGO_DB")).Collection("users")
	if insertResult, err := collection.InsertOne(context.TODO(), user); err == nil {
		if oid, ok := insertResult.InsertedID.(primitive.ObjectID); ok {
			user.ID = oid
		}
	} else {
		log.Fatal(err)
	}

	b, err := json.Marshal(user)
	if err != nil {
		log.Println(err)
		return
	}

	w.Write([]byte(string(b)))
}
