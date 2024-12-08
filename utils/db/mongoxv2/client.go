package mongoxv2

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var client *mongo.Client // default client
var db Database          // default database

type Database struct {
	*mongo.Database
}

// Connect creates a new client and connects to the given MongoDB
// instance at the given URI and database name.
//
// It returns an error if the connection cannot be established.
func Connect(ctx context.Context, uri string, dbName string) error {
	var err error
	// Create a new client with the given URI.
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return err
	}
	// Set the default database of the client to the given database name.
	db.Database = client.Database(dbName)

	// Check the connection by sending a ping message to the server.
	// If an error occurs, return it.
	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	return nil
}

// SetDatabase sets the default database of the client to the given database.
//
// This will override the database set in Connect.
func SetDatabase(database *mongo.Database) {
	db.Database = database
	client = db.Client()
}

func GetClient() *mongo.Client {
	return client
}

// CollRead returns a collection with the given name from the default database,
// configured for nearest read preference.
//
// This is useful for reads that do not require consistency across the entire
// replica set, such as when loading a user's profile information.
func CollRead(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Nearest()))
}

// CollWrite returns a collection with the given name from the default database,
// configured for primary read preference.
//
// This is useful for writes that require consistency across the entire replica set,
// such as when creating a new user.
func CollWrite(collName string) *mongo.Collection {
	return db.Collection(collName, options.Collection().SetReadPreference(readpref.Primary()))
}
