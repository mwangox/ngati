package service

import (
	"github.com/go-resty/resty/v2"
	"log"
	"nitra/registry/model"
	"nitra/registry/utils/propertymanager"
	"nitra/registry/utils/redis"
	"time"
)

func HealthCheckHandler()  {

	for{
		serviceMap, err := redis.HGetAll("services")
		if err != nil {
			log.Println("Failed to retrieve services:" ,err)
			time.Sleep(time.Duration(propertymanager.GetIntProperty("health-check.retry.delay-on-fail")) * time.Second)
			HealthCheckHandler()
		}

		if len(serviceMap) == 0 {
			log.Println("No services found")
			time.Sleep(time.Duration(propertymanager.GetIntProperty("health-check.retry.delay-on-no-service")) * time.Second)
			HealthCheckHandler()
		}
		for serviceName, serviceJson  := range serviceMap{
			serviceDef, err := model.FromJson(serviceJson)
			if err != nil{
				log.Println("Failed to decode service: ", serviceName)
			}
			go check(serviceName, serviceDef)
		}
		time.Sleep(time.Duration(propertymanager.GetIntProperty("health-check.interval")) * time.Second)
	}
}

func check(serviceName string, service *model.ServiceDefinition)  {
	healthCheckUri := service.HttpScheme + "://" + service.ServiceIp + ":" + service.ServicePort + service.HealthCheckPath
	client := resty.New().R()
	response, err := client.Get(healthCheckUri)
	log.Println("Health check status response:", response, serviceName)

	service.Status = "OK"
	if err != nil || !response.IsSuccess() {
		log.Println("Health check for ", serviceName, " failed", err)
		service.Status = "NOT_OK"
	}

	updateStatus(serviceName, service.ToJson())
}

func updateStatus(serviceName, serviceJson string)  {
	if err := redis.HSet("services", serviceName, serviceJson); err != nil{
		log.Println("Failed to update service status:", err)
	}
}