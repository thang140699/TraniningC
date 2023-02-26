package main

import (
	"crawl/database/mongo"
	database "crawl/database/mongo"
	"crawl/database/repository"
)

type Provider struct {
	*mongo.MongoProvider
}

type Container struct {
	*Provider
	Config                  Config
	MalshareDailyRepository repository.MalshareDailyRepository
}

func NewContainer(config Config) (*Container, error) {
	container := new(Container)
	err := container.InitContainer(config)
	if err != nil {
		return nil, err
	}

	container.Config = config

	return container, nil
}
func (container *Container) InitContainer(config Config) error {
	// Load providers into container
	err := container.LoadProviders(config)
	if err != nil {
		return err
	}

	// Load repositories
	container.LoadRepositoryImplementations(config)

	return nil
}
func (container *Container) LoadProviders(config Config) error {

	mongoProvider := mongo.NewMalshareMongoRepository(config.MongoURL)

	container.Provider = &Provider{
		MongoProvider: mongoProvider,
	}
	return nil
}

func (container *Container) LoadRepositoryImplementations(config Config) {
	container.MalshareDailyRepository = database.NewMalshareMongoRepository(container.MongoProvider)

}
