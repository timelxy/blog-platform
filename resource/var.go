package resource

import (
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoClient
var MongoClient *mongo.Client

// Validater
var Validate *validator.Validate
