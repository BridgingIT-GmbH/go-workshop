package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
)

func CreateClient(mongoUri string) *mongo.Client {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUri))
	if err != nil {
		log.Panic(fmt.Sprintf("Could not connect to MongoDB! URI: '%v'", mongoUri))
	}
	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		log.Panic(fmt.Sprintf("MongoDB ist not reachable! URI: '%v'", mongoUri))
	}
	log.Print(fmt.Sprintf("Connected to MongoDB: '%v'", mongoUri))
	return client
}
