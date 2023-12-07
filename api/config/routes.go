package config

import (
	"checkout/api/controllers"
)

func init() {
	v1 := engine.Group("/v1")
	{
		v1.GET("/test", controllers.GetTest)
	}
	engine.GET("/health/check", controllers.HealthCheck)
}
