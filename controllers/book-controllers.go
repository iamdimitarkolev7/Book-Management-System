package controllers

import (
	"book-management-system/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection

// creating the connection to MongoDB
// init function is called once, when we run the project
func init() {
	loadTheEnv()
	createDBinstance()
}

func loadTheEnv() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading the .env file")
	}
}

func createDBinstance() {
	connectionString := os.Getenv("DB_URI")
	dbName := os.Getenv("DB_NAME")
	collectionName := os.Getenv("DB_COLLECTION_NAME")

	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// connect to mongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection instance created!")
}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	payload := getAllBooks()
	json.NewEncoder(w).Encode(payload)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	payload := getBookById(params["id"])
	json.NewEncoder(w).Encode(payload)
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var book models.Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	createBook(book)
	json.NewEncoder(w).Encode(&book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	var updatedBook models.Book
	_ = json.NewDecoder(r.Body).Decode(&updatedBook)
	var params = mux.Vars(r)
	updateBook(params["id"], updatedBook)
	json.NewEncoder(w).Encode(&updatedBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	params := mux.Vars(r)
	deleteBook(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func getAllBooks() []primitive.M {
	curr, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	var results []primitive.M

	for curr.Next(context.Background()) {
		var result bson.M
		err := curr.Decode(&result)

		if err != nil {
			log.Fatal(err)
		}

		results = append(results, result)
	}

	if err := curr.Err(); err != nil {
		log.Fatal(err)
	}

	curr.Close(context.Background())

	return results
}

func getBookById(bookId string) primitive.M {
	id, _ := primitive.ObjectIDFromHex(bookId)
	result := bson.M{"_id": id}

	err := collection.FindOne(context.Background(), result).Decode(&result)

	if err != nil {
		log.Fatal(err)
	}

	return result
}

func createBook(book models.Book) {
	createdBook, err := collection.InsertOne(context.Background(), book)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("A new book is successfully created with id: ", createdBook.InsertedID)
}

func updateBook(bookId string, updatedBook models.Book) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"title": updatedBook.Title, "author": updatedBook.Author}}

	_, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully updated book")
}

func deleteBook(bookId string) {
	id, _ := primitive.ObjectIDFromHex(bookId)
	bookToDelete := bson.M{"_id": id}

	_, err := collection.DeleteOne(context.Background(), bookToDelete)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully deleted a book")
}
