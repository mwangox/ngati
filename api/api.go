package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net"
	"nitra/registry/controller"
	"nitra/registry/utils/propertymanager"
)

func InitializeAPIs() error {
	router := GetRoutes()
	router.Use(cors.Default())
	return router.Run(net.JoinHostPort("", propertymanager.GetStringProperty("registry.port")))
}

func GetRoutes() *gin.Engine  {
	r := gin.Default()

	r.POST("/ngati/services/register/:serviceName", controller.RegisterHandler)
	r.DELETE("/ngati/services/de-register/:serviceName", controller.DeRegisterHandler)
	r.GET("/ngati/services/:serviceName", controller.QueryHandler)
	r.GET("/ngati/services", controller.QueryAllHandler)
	r.PUT("/ngati/services/:serviceName/metadata", controller.UpdateHandler)
	r.GET("/health", controller.HealthHandler)

	//KV handlers
	r.GET("/ngati/kv/:keyName", controller.GetKeyHandler)
	r.GET("/ngati/kv", controller.GetAllKeysHandler)
	r.PUT("/ngati/kv", controller.SetKeysHandler)
	return r
}
