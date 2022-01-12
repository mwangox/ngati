package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/tidwall/sjson"
	"log"
	"net/http"
	"nitra/registry/model"
	"nitra/registry/utils/redis"
)

func UpdateHandler(ctx *gin.Context) {
	serviceName := ctx.Param("serviceName")

	serviceJson, err := redis.HGet("services", serviceName)
	if err != nil {
		if redis.IsRedisNil(err){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004,"Service not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1005, "Failed to retrieve service details"))
		return
	}

	query := ctx.Request.URL.Query()
    newServiceJson := serviceJson
	for key, value := range query{
		if !isValidField(key){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1009, "Service metadata  " + key + " is not defined"))
			return
		}
		var err error
		newServiceJson, err = sjson.Set(newServiceJson, key, value[0])
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1007,"Failed to update service param: " + key))
			return
		}
	}

	if err := redis.HSet("services", serviceName, newServiceJson); err != nil{
		log.Println("Failed to update service metadata:", err)
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1008, "Failed to update service metadata"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Success",
	})
}

func isValidField(field string) bool  {
	metadata := []string{ "service_name", "service_port", "service_ip", "service_dc", "health_check_path"}
	for _, value := range metadata{
		if value == field{
			return true
			break
		}
	}
	return false
}