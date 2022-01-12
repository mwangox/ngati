package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HealthHandler(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"status":"OK",
	})
}
