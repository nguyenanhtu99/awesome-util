package mongox

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var MongoClient *mongo.Client

func ConnectMongo(ctx context.Context) error {
	opts := options.Client().ApplyURI("mongodb://localhost:27017")

	mgmConf := mgm.Config{CtxTimeout: time.Second * 10} //nolint:mnd // 10 seconds is enough for timeout
	err := mgm.SetDefaultConfig(&mgmConf, "mydb", opts)
	if err != nil {
		return fmt.Errorf("failed to set MongoDB default config: %w", err)
	}

	_, client, _, err := mgm.DefaultConfigs()
	if err != nil {
		return fmt.Errorf("failed to get client, err: %w", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return fmt.Errorf("failed to ping to MongoDB: %w", err)
	}

	log.Println("Connected to MongoDB!")
	return nil
}
