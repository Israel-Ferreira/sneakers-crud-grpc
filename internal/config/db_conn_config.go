package config

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectDb(host, username, port, password string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	dbConnectionStr := fmt.Sprintf("mongodb://%s:%s@%s:%s", username, password, host, port)

	connection, err := mongo.Connect(ctx, options.Client().ApplyURI(dbConnectionStr))

	if err != nil {
		log.Fatalln(err.Error())
	}

	MongoClient = connection
}
