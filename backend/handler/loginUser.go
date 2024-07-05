package handler

import (
	"net/http"

	"github.com/shh4und/movie-tracker/auth"
	"github.com/shh4und/movie-tracker/config"

	"github.com/gin-gonic/gin"
	"github.com/shh4und/movie-tracker/schemas"
)

func LoginUser(ctx *gin.Context) {
	request := LoginUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("request validation error: %v", err)
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	user := schemas.User{}

	if err := db.Where("username = ?", request.Username).First(&user).Error; err != nil {
		sendError(ctx, http.StatusBadRequest, "invalid username or password")
		return
	}

	if !auth.ComparePasswords(user.Password, []byte(request.Password)) {
		sendError(ctx, http.StatusBadRequest, "invalid username or password")
		return
	}

	secret := []byte(config.Envs.JwtToken)
	token, err := auth.CreateJWT(secret, user.ID)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	sendSuccess(ctx, "login-user", token)

}
