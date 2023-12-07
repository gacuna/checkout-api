package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
	Status string `json:"status"`
}

func healthCheck(context *gin.Context) {
	var health = Health{
		Status: "OK",
	}
	context.IndentedJSON(http.StatusOK, health)
}

func getTest(context *gin.Context) {
	context.JSON(http.StatusOK, "test")
}

func main() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.GET("/test", getTest)
	}
	router.GET("/health/check", healthCheck)
	http.ListenAndServe(":9090", router)
}
