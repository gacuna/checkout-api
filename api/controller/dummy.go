package controller

import (
	"github.com/gin-gonic/gin"
	//	"github.com/jinzhu/gorm"
	//	"gin-restful-best-practice/middlewares"
	//	"gin-restful-best-practice/models"

	"net/http"
)

func GetTest(context *gin.Context) {
	context.JSON(http.StatusOK, "test")
}
