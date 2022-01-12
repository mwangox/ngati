package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nitra/registry/model"
	"nitra/registry/utils/redis"
)

func QueryHandler(ctx *gin.Context) {

	serviceName := ctx.Param("serviceName")

	serviceJson, err := redis.HGet("services", serviceName)
	if err != nil {
		if redis.IsRedisNil(err){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004,"Service not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1005,"Failed to retrieve service details"))
		return
	}

	serviceDef, err := model.FromJson(serviceJson)
	if err != nil{
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1006,"Generic processing failure"))
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Success",
		"body": serviceDef,
	})
}


func QueryAllHandler(ctx *gin.Context) {

	serviceMap, err := redis.HGetAll("services")
	if err != nil {
		if redis.IsRedisNil(err){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004,"No services found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1005, "Failed to retrieve service details"))
		return
	}

	if len(serviceMap) == 0 {
		ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004,"No services found"))
		return
	}

	newServiceMap := make(map[string]*model.ServiceDefinition)
	for serviceName, serviceJson  := range serviceMap{
		serviceDef, err := model.FromJson(serviceJson)
		if err != nil{
			ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1006,"Generic processing failure"))
			return
		}
		newServiceMap[serviceName] = serviceDef
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Success",
		"body": newServiceMap,
	})
}