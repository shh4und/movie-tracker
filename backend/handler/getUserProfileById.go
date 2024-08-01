package handler

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// 	"github.com/shh4und/movie-tracker/schemas"
// )

// func GetUserProfileByID(ctx *gin.Context) {
// 	id := ctx.Query("id")

// 	if id == "" {
// 		sendError(ctx, http.StatusBadRequest, errParamIsRequired("id", "query-param").Error())
// 		return
// 	}
// 	user := schemas.User{}

// 	if err := db.First(&user, id).Error; err != nil {
// 		sendError(ctx, http.StatusNotFound, fmt.Sprintf("user with id: %s not found on the database", id))
// 		return
// 	}

// 	sendSuccess(ctx, "get-user-id", user)

// }
