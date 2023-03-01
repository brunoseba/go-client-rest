package config

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func DBconn(ctx context.Context) mongo.Client {
	// connect with database --------------------------------------------------------
	mongoconn := options.Client().ApplyURI("mongodb://127.0.0.1:27017")
	mongoclient, err := mongo.Connect(ctx, mongoconn)
	if err != nil {
		log.Fatal(err)
	}

	if err := mongoclient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Mongodb connection successful")
	// -------------------------------------------------------------------------------

	return *mongoclient

}
