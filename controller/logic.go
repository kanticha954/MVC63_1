package controller

import (
	"MVC63_1/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListHospital(c *gin.Context) {
	//call pugin
	data, er := model.GetHospital()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
