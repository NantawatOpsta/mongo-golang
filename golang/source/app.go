package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Book struct {
	Title      string    `bson:"title,omitempty"`
	Created_at time.Time `bson:"created_at"`
}

func main() {
	fmt.Println("run main function")
	bson_collection()
}

func bson_collection() {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://root:root@mongodb:27017/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	collection := client.Database("golang").Collection("bson_db")

	user := bson.D{{"title", "The Room"}, {"type", "English Breakfast"}, {"rating", 6}, {"created_at", time.Now()}}

	result, err := collection.InsertOne(context.TODO(), user)

	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	fmt.Println("completed insert bson record")

	var books []Book
	cursor, err := collection.Find(context.TODO(), bson.D{{"title", "The Room"}})
	cursor.All(ctx, &books)

	for i := range books {
		fmt.Println(books[i].Title)
		fmt.Println(books[i].Created_at.Format("2006-02-01"))
	}

}
