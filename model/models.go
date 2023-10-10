package model

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type Models interface {
	getCollection() string
	getDB() string
	GetReadDB() (*mongo.Collection, error)
	GetWriteDB() (*mongo.Collection, error)
}
