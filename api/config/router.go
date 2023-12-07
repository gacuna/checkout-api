package config

import (
	"github.com/gin-gonic/gin"
)

var engine = gin.Default()

func Create_engine() *gin.Engine {
	return engine
}
