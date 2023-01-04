package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Post struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Text string             `json:"text" bson:"text"`
	Imgs []string           `json:"imgs" bson:"imgs"`
}
