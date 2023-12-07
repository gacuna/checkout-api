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

func main() {
	router := gin.Default()
	router.GET("/health/check", healthCheck)
	router.Run("localhost:9090")
}
