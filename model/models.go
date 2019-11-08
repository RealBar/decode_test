package model

import (
	"decode_test/pkg/config"
	"decode_test/pkg/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func Setup(){
	// init mongo client
	mongoURI := 
	timeout := utils.DBConnectTimeout
	options.Client().ApplyURI()
}