package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var SendUserinfo = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}

var GetUserinfo = func(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, nil)
}