package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nitra/registry/model"
	"nitra/registry/utils/redis"
)

func RegisterHandler(ctx *gin.Context) {
		service := &model.ServiceDefinition{}
		if err := ctx.ShouldBindJSON(service); err != nil{
			ctx.JSON(http.StatusBadRequest, model.NewApiErrorResponse(1001, "Failed to decode registration request"))
			return
		}

		log.Println("Registration request received is:", service.ToJson())
		serviceName := ctx.Param("serviceName")
		log.Println("ServiceName is:", serviceName)
		if err := redis.HSet("services", ctx.Param("serviceName"), service.ToJson()); err != nil{
			log.Println("Registration process failed:", err)
			ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1002, "Failed perform service registration"))
			return
		}

		ctx.JSON(http.StatusCreated, gin.H{
			"status_code":"1000",
			"status_description":"Successfully registered",
		})
}
