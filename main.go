package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-mongo-implementation/api"
	"github.com/go-mongo-implementation/pkg/config"
	"github.com/go-mongo-implementation/pkg/mongo"
	"github.com/go-mongo-implementation/repository"
)

const (
	TIMEOUT       = 10
	DATABASE_NAME = "test"
)

var (
	cfg *config.AppConfig
)

type routes struct {
	router *gin.Engine
}

func main() {
	cfg = config.GetConfig()

	var mongoDbClient, err = mongo.NewClient(cfg.MongoDB, TIMEOUT*time.Second)
	if err != nil {
		panic("MongoDB connection timeout")
	}

	var userRepository = repository.NewUserRepository(mongoDbClient, DATABASE_NAME)

	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", func(ctx *gin.Context) {
			api.GetUsers(ctx, userRepository)
		})
		v1.POST("/user", func(ctx *gin.Context) {
			api.InsertUser(ctx, userRepository)
		})
	}

	router.Run(":1453")
}
