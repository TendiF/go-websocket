package utils

import (
	"context"
	"log"
	"time"
	"os"
	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/mongo/readpref"
)

type Post struct {

	Title string `json:"title,omitempty"`
	
	Body string `json:"body,omitempty"`
	
}

func CreateConnection(){
	log.Println(os.Getenv("MONGO_DB"))
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://"+os.Getenv("MONGO_USERNAME")+":"+os.Getenv("MONGO_PASSWORD")+"@"+os.Getenv("MONGO_HOSTNAME")+":"+os.Getenv("MONGO_PORT")+"/?authSource=admin"))
	if err != nil {
			log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	post := Post{"title", "Post body"}
	collection := client.Database(os.Getenv("MONGO_DB")).Collection("posts")
	insertResult, err := collection.InsertOne(context.TODO(), post)

	if err != nil {

	log.Fatal(err)

	}



	log.Println("Inserted post with ID:", insertResult.InsertedID)
}
