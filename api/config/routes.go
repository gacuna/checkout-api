package config

import (
	"checkout/api/controller"
)

func init() {
	v1 := engine.Group("/v1")
	{
		v1.GET("/test", controller.GetTest)
		v1.GET("/checkout", controller.GetTest)
		v1.POST("/checkout", controller.PostTest)
	}
	engine.GET("/health/check", controller.HealthCheck)
}
