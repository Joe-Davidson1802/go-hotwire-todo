package main

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {
	ID       primitive.ObjectID `bson:"_id"`
	Title    string             `bson:"title"`
	Complete bool               `bson:"complete"`
}
