package mongo

import (
	"fmt"
	"time"

	"github.com/globalsign/mgo"
	"github.com/go-mongo-implementation/pkg/config"
)

type Client struct {
	Session *mgo.Session
}

type MongoDBClient interface {
	NewSession() *mgo.Session
	CreateIndex(fields []string, isUnique bool, indexName, databaseName, collectionName string) error
}

func NewClient(configuration config.MongoDBConfig, timeout time.Duration) (MongoDBClient, error) {
	connectionString := fmt.Sprintf("mongodb://%s:%s@%s", configuration.Username, configuration.Password, configuration.Host)
	session, err := mgo.DialWithTimeout(connectionString, timeout)

	if err != nil {
		return nil, err
	}

	return &Client{Session: session}, nil
}

func (c *Client) NewSession() *mgo.Session {
	newSession := c.Session.Copy()
	newSession.SetMode(mgo.Strong, true)
	return newSession
}

func (c *Client) CreateIndex(fields []string, isUnique bool, indexName, databaseName, collectionName string) error {
	localSession := c.NewSession()
	defer localSession.Close()

	index := mgo.Index{
		Key:        fields,
		Unique:     isUnique,
		Name:       indexName,
		Background: true,
	}

	col := localSession.DB(databaseName).C(collectionName)

	if err := col.EnsureIndex(index); err != nil {
		return err
	}

	return nil
}
