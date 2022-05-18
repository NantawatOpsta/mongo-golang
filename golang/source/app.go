package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	var ctx = context.TODO()

	clientOptions := options.Client().ApplyURI("mongodb://root:root@mongodb:27017/")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connected to MongoDB")
	}

	// insert database and collection with date time
	collection := client.Database("golang").Collection("golang")

	// insert data title, content, created_at
	user := []interface{}{
		bson.D{{"type", "English Breakfast"}, {"rating", 6}},
		bson.D{{"type", "Oolong"}, {"rating", 7}, {"vendor", bson.A{"C"}}},
		bson.D{{"type", "Assam"}, {"rating", 5}},
		bson.D{{"type", "Earl Grey"}, {"rating", 8}, {"vendor", bson.A{"A", "B"}}},
	}

	result, err := collection.InsertMany(context.TODO(), user)

	if err != nil {
		panic(err)
	}
	// display the id of the newly inserted object
	fmt.Println(result)
}
