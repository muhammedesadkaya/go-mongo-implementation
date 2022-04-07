package api

import (
	"github.com/gin-gonic/gin"
	"github.com/go-mongo-implementation/model"
	"github.com/go-mongo-implementation/repository"
)

func GetUsers(ctx *gin.Context, userRepository repository.UserRepository) {
	var (
		users []*model.User
		err   error
	)

	if users, err = userRepository.GetAll(ctx); err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(200, users)
}

func InsertUser(ctx *gin.Context, userRepository repository.UserRepository) {
	var (
		user *model.User
		err  error
	)

	if err = ctx.BindJSON(&user); err != nil {
		ctx.JSON(500, err)
	}

	newUser := model.NewUser(user.FullName, user.Identity, user.PhoneNumber, user.Email, user.Gender)

	if err = userRepository.Insert(ctx, newUser); err != nil {
		ctx.JSON(500, err)
	}

	ctx.JSON(200, nil)
}
