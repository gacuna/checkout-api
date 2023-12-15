package controller

import (
	"checkout/api/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostTest(context *gin.Context) {
	var newDummy model.Dummy
	if err := context.BindJSON(&newDummy); err != nil {
		return
	}
	context.IndentedJSON(http.StatusCreated, newDummy)
}
