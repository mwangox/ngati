package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"nitra/registry/model"
	"nitra/registry/utils/redis"
)

func GetKeyHandler(ctx *gin.Context) {
	key := ctx.Param("keyName")
	value, err := redis.HGet("kv_store", key)
	if err != nil{
		if redis.IsRedisNil(err){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004, "Key not found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1005, "Failed to retrieve key value"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Success",
		key: value,
	})
}

func SetKeysHandler(ctx *gin.Context) {
	query := ctx.Request.URL.Query()
	for key, value := range query{
		if err := redis.HSet("kv_store", key, value[0]); err != nil{
			log.Println("Failed to set value for a key:", err)
			ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1012, "Failed to set value for a key"))
			return
		}
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"status_code":"1000",
		"status_description":"Success",
	})
}

func GetAllKeysHandler(ctx *gin.Context)  {
	kvStore, err := redis.HGetAll("kv_store")
	if err != nil {
		if redis.IsRedisNil(err){
			ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004, "No keys found"))
			return
		}
		ctx.JSON(http.StatusInternalServerError, model.NewApiErrorResponse(1005, "Failed to retrieve all keys"))
		return
	}

	if len(kvStore) == 0 {
		ctx.JSON(http.StatusNotFound, model.NewApiErrorResponse(1004,"No keys found"))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status_code":"1000",
		"status_description":"Success",
		"body": kvStore,
	})
}