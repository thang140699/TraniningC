package provider

import (
	download "crawl/utilities"
	"fmt"
	"log"
)

type MongoProvider struct {
	mongoClient *download.MongoDB
}

func NewMongoProviderFromURL(u string) *MongoProvider {
	client := download.NewMongoDBFromURL(u)
	if client == nil {
		log.Fatalf("Mongo server connected unsuccessfully %s", u)
	}

	return &MongoProvider{
		mongoClient: client,
	}
}
func NewMongoProvider(server, database, collection string) *MongoProvider {
	configMongo := make(map[string]string)
	configMongo[download.DB_SERVER] = server
	configMongo[download.DB_DATABASE] = database
	configMongo[download.DB_COLLECTION] = collection

	client := download.NewMongoDB(configMongo)
	if client == nil {
		log.Fatalf("Mongo server connected unsuccessfully: nill client")
	}

	return &MongoProvider{
		mongoClient: client,
	}
}

func (provider *MongoProvider) MongoClient() *download.MongoDB {
	return provider.mongoClient
}

func (provider *MongoProvider) NewError(e error) error {
	if e == nil {
		return nil
	}
	return DatabaseExecutionError{
		Err:     e,
		Message: fmt.Sprintf("Mongo execution error: %s", e.Error()),
	}
}

type DatabaseExecutionError struct {
	Err     error
	Message string
}

func (e DatabaseExecutionError) Error() string {
	return e.Message
}

func (e DatabaseExecutionError) Unwrap() error {
	return e.Err
}
