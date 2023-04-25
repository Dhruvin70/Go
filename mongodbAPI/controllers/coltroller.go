package controller

import (
	"context"
	"fmt"
	"log"

	modaleStruct "mongoAPI/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectstring = "mongodb+srv://Dhruvin:So29vFlfmjgjtD0F@cluster1.1yrdco1.mongodb.net/?retryWrites=true&w=majority"
const dbName = "Netflix"
const collectionName = "watchlist"

// IMP deploy connection

var collection *mongo.Collection

// connect with mongoDB
// init --> runs first and only once(inbuilt)

func init() {
	// client options
	// connecting with any DB

	clientOption := options.Client().ApplyURI(connectstring)

	// connect to mongodb
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB connection success")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("Collection reference is ready")

}

func InsertOneMoie(movie modaleStruct.Netflix) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted one value in db with id: ", inserted.InsertedID)
}

func UpdateOnemovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("uptated count", result.ModifiedCount)

}

// delete one record

func DeleteOneMovie(movieId string) {

	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	deletecount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Movies deleated", deletecount)
}

// delete all record from mongoDB

func DeleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("deleted count", deleteResult.DeletedCount)

	return deleteResult.DeletedCount
}

// get all movies from database

func GetAllMovies() []primitive.M {
	cursour, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var Movies []primitive.M

	for cursour.Next(context.Background()) {
		var movie bson.M
		err := cursour.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		Movies = append(Movies, movie)
	}

	defer cursour.Close(context.Background())
	return Movies
}
