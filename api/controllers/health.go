package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
	Status string `json:"status"`
}

func HealthCheck(context *gin.Context) {
	var health = Health{
		Status: "OK",
	}
	context.IndentedJSON(http.StatusOK, health)
}
