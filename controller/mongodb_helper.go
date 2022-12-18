package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/hiteshshimpi-55/go-mongo-api/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

/*
	MongoDB Helper Methods
	- Get All Movies
	- Get One Movie
	- Update One Movie
	- Delete One Movie
	- Delete All Movies
*/

func getAllMovies() []primitive.M {

	// find the document from collection... here bson.D{{params}}. it is empty means retrive all docs

	docs, err := collection.Find(context.Background(), bson.D{{}})

	if err != nil {
		log.Fatal(err)
	}

	// slice containing all movies from docs
	var movies []primitive.M

	// Looping through the documents
	for docs.Next(context.Background()) {

		var movie bson.M
		err := docs.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}

		movies = append(movies, movie) // append the movie in slice
	}
	defer docs.Close(context.Background())
	return movies
}

func getOneMovie(movie_title string) models.Movie {

	// id, _ := primitive.ObjectIDFromHex(movie_id)
	filter := bson.M{"title": movie_title} //condition || Filter for searching document

	var res models.Movie
	err := collection.FindOne(context.Background(), filter).Decode(&res)

	if err != nil {
		log.Fatal(err)
	}

	return res
}

func setWatched(movie_id string) int64 {
	id, _ := primitive.ObjectIDFromHex(movie_id)
	filter := bson.M{"_id": id}

	update := bson.M{"$set": bson.M{"watched": "true"}}

	result, err := collection.UpdateOne(context.Background(), filter, update)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
	return result.MatchedCount
}

func deletOneMovie(movie_id string) int64 {

	id, _ := primitive.ObjectIDFromHex(movie_id)
	filter := bson.M{"_id": id}

	res, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	return res.DeletedCount
}

func deletAllMovies() int64 {

	res, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%d Items got deleted from database", res.DeletedCount)
	return res.DeletedCount
}

func insertMovie(movie models.Movie) {
	res, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted new movie in database\n", res.InsertedID)
}
