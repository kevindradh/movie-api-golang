package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Movie struct {
	ID       primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty" swaggerignore:"true"`
	Title    string             `json:"title" bson:"title" binding:"required"`
	Director string             `json:"director" bson:"director" binding:"required"`
	Year     int                `json:"year" bson:"year" binding:"required"`
	Genre    string             `json:"genre" bson:"genre"`
	Rating   float64            `json:"rating" bson:"rating" binding:"min=0,max=10"`
	Poster   string             `json:"poster" bson:"poster"`
}
