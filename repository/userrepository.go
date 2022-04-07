package repository

import (
	"context"
	"fmt"

	"github.com/go-mongo-implementation/model"
	"github.com/go-mongo-implementation/pkg/mongo"
)

const (
	COLLECTION_NAME = "user"
)

type UserRepository interface {
	GetAll(ctx context.Context) ([]*model.User, error)
	Insert(ctx context.Context, user *model.User) error
}

type userRepository struct {
	mongoClient  mongo.MongoDBClient
	databaseName string
}

func (self *userRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	var records []*model.User

	err := session.
		DB(self.databaseName).
		C(COLLECTION_NAME).
		Find(nil).All(&records)

	if err != nil {
		return nil, err
	}

	return records, nil
}

func (self *userRepository) Insert(ctx context.Context, user *model.User) error {
	var session = self.mongoClient.NewSession()
	defer session.Close()

	err := session.
		DB(self.databaseName).
		C(COLLECTION_NAME).
		Insert(user)

	if err != nil {
		return err
	}

	return nil
}

func NewUserRepository(client mongo.MongoDBClient, databaseName string) UserRepository {

	if err := client.CreateIndex([]string{"CreateDate", "FullName"}, true, "CreateDate_FullName", databaseName, COLLECTION_NAME); err != nil {
		fmt.Println(err)
	}

	return &userRepository{mongoClient: client, databaseName: databaseName}
}
