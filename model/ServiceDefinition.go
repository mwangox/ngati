package model

import "encoding/json"

type ServiceDefinition struct {
	ServiceName     string `json:"service_name"`
	ServiceIp       string `json:"service_ip"`
	ServicePort     string `json:"service_port"`
	ServiceDc       string `json:"service_dc"`
	HttpScheme      string `json:"http_scheme"`
	HealthCheckPath string `json:"health_check_path"`
	Status          string `json:"status"`
}

func(t *ServiceDefinition) ToJson() string {
	jsonBytes, _ := json.Marshal(t)
	return string(jsonBytes)
}

func FromJson(service string) (*ServiceDefinition, error) {
	serviceObject := &ServiceDefinition{}
	err := json.Unmarshal([]byte(service), serviceObject)
	return serviceObject, err
}