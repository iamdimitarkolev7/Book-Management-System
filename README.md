# Book-Management-System
Creating a simple REST API using Golang and MongoDB

## Purpose of the project
The purpose of the project is to create a database management system to keep track of books. Possible actions are: displaying all the books, display a particular book by a given id, creating a new book, editing an existing book, deleting an existing book.

## How to use the project
You can use this REST API as a server for your client application.

## Possible routes
* GET request to this route _/api/books_ will display all the books in the database
* GET request to this route _/api/books/{id}_ will display a single book by its id
* POST request to this route _/api/books/create_ will add a new book to the database
* PUT request to this route _/api/books/update/{id}_ will update a book with the given _id_
* DELETE request to this route _/api/books/delete/{id}_ will delete a book with the given _id_

## DB Model
The database model is a Book, which has a unique id, name and author
