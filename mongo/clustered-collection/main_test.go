package main

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	uri                  = "mongodb://localhost:27017"
	dbName               = "performance_test"
	clusteredCollName    = "clustered_logs"
	nonClusteredCollName = "non_clustered_logs"
	client               *mongo.Client
	db                   *mongo.Database
	numDocuments         = 1_000_000
)

func setup() {
	if client == nil {
		var err error
		client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
		if err != nil {
			log.Fatal(err)
		}
		db = client.Database(dbName)
	}
}

func BenchmarkQueryClusteredCollection(b *testing.B) {
	setup()

	// Drop collection before the benchmark to ensure fresh test data
	db.Collection(clusteredCollName).Drop(context.TODO())

	// // Create Clustered Collection
	opts := options.CreateCollection().
		SetClusteredIndex(bson.D{{Key: "key", Value: bson.D{{Key: "_id", Value: 1}}}, {Key: "unique", Value: true}})

	err := db.CreateCollection(context.TODO(), clusteredCollName, opts)
	if err != nil {
		log.Fatal("Error creating clustered collection: ", err)
	}

	coll := db.Collection(clusteredCollName)
	startTime := time.Now()
	insertTestData(coll)
	log.Printf("Clustered Collection: Inserted %d documents in %s\n", numDocuments, time.Since(startTime))

	startID := primitive.NewObjectIDFromTimestamp(startTime)
	endID := primitive.NewObjectIDFromTimestamp(time.Now())

	b.ResetTimer() // Reset the timer for the actual query benchmarking

	for i := 0; i < b.N; i++ {
		filter := bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "$gte", Value: startID},
				{Key: "$lt", Value: endID},
			}},
		}

		_, err := coll.Find(context.TODO(), filter)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkQueryNonClusteredCollection(b *testing.B) {
	setup()

	// Drop collection before the benchmark to ensure fresh test data
	db.Collection(nonClusteredCollName).Drop(context.TODO())

	// Create Non-Clustered Collection and Index
	err := db.CreateCollection(context.TODO(), nonClusteredCollName)
	if err != nil {
		log.Fatal("Error creating non-clustered collection: ", err)
	}

	coll := db.Collection(nonClusteredCollName)
	startTime := time.Now()
	insertTestData(coll)
	log.Printf("Non-Clustered Collection: Inserted %d documents in %s\n", numDocuments, time.Since(startTime))

	startID := primitive.NewObjectIDFromTimestamp(startTime)
	endID := primitive.NewObjectIDFromTimestamp(time.Now())

	b.ResetTimer() // Reset the timer for the actual query benchmarking

	for i := 0; i < b.N; i++ {
		filter := bson.D{
			{Key: "_id", Value: bson.D{
				{Key: "$gte", Value: startID},
				{Key: "$lt", Value: endID},
			}},
		}

		_, err := coll.Find(context.TODO(), filter)
		if err != nil {
			b.Fatal(err)
		}
	}
}

func insertTestData(coll *mongo.Collection) {
	for n := 0; n < numDocuments/100; n++ {
		docs := make([]interface{}, 100)
		for i := 0; i < 100; i++ {
			docs[i] = bson.D{
				{Key: "log", Value: fmt.Sprintf("Log entry %d", n*100+i)},
			}
		}
		_, err := coll.InsertMany(context.TODO(), docs)
		if err != nil {
			log.Fatal(err)
		}
	}
}
