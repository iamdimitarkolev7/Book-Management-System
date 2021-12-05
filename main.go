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

func main() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb+srv://kolev7:kolev7@book-management-cluster.1ahta.mongodb.net/myFirstDatabase?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	bookManagerDatabase := client.Database("book-manager")
	booksCollection := bookManagerDatabase.Collection("books")

	bookResult, err := booksCollection.InsertMany(ctx, []interface{}{
		bson.D{
			{Key: "Name", Value: "Harry Potter"},
			{Key: "Author", Value: "J.K.Rowling"},
		},
		bson.D{
			{Key: "Name", Value: "Golang Intro"},
			{Key: "Author", Value: "Dimitar Kolev"},
		},
	})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Inserted %v books\n", len(bookResult.InsertedIDs))
}
