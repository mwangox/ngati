package cmd

import (
	"log"
	"nitra/registry/api"
	"nitra/registry/service"
	"nitra/registry/utils/propertymanager"
	"nitra/registry/utils/redis"
)

func Start() error {
	log.Println("Initialize redis client...")
	redis.InitializeRedis()

	log.Println("Initialize configurations...")
	if err := propertymanager.InitializeConfig(); err != nil{
		return err
	}

	log.Println("Start async health check monitoring...")
	go service.HealthCheckHandler()

	log.Println("Start the registry web server...")
	return api.InitializeAPIs()
}
