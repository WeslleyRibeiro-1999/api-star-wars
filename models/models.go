package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Planet struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Nome    string             `json:"nome" bson:"nome"`
	Clima   string             `json:"clima" bson:"clima"`
	Terreno string             `json:"terreno" bson:"terreno"`
	Qtde    int32              `json:"qtde_filmes,omitempty" bson:"qtde_filmes,omitempty"`
}
