package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshshimpi-55/go-mongo-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connection string = "mongodb://localhost:27017"

const db_name = "Netflix"

const collection_name = "favourites"

var collection *mongo.Collection

func init() {

	//Client Option
	client_options := options.Client().ApplyURI(connection)

	//Connection to mongo
	client, err := mongo.Connect(context.TODO(), client_options)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection Successfull with MongoDB database")

	collection = client.Database(db_name).Collection(collection_name)

	fmt.Println("Collection instance ready. . . ")

}

// MongoDB controllers

func GetAllMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	all_movies := getAllMovies()
	json.NewEncoder(w).Encode(all_movies)
}

func GetOneMovies(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	res := getOneMovie(params["id"])
	json.NewEncoder(w).Encode(res)
}

func Update(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")
	params := mux.Vars(r)
	res := setWatched(params["id"])
	json.NewEncoder(w).Encode(res)
}

func DeleteOneMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	params := mux.Vars(r)
	res := deletOneMovie(params["id"])
	json.NewEncoder(w).Encode(res)
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")
	res := deletAllMovies()
	json.NewEncoder(w).Encode(res)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")
	var item models.Movie
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		log.Fatal(err)
	}
	insertMovie(item)
	json.NewEncoder(w).Encode(item)
}
