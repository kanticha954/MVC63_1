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

func ListPatient(c *gin.Context) {
	//call pugin
	data, er := model.GetPatient()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func ListPatientStatus(c *gin.Context) {
	//call pugin
	data, er := model.GetPatientStatus()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func CountPatient(c *gin.Context) {
	//call pugin
	data, er := model.CountPatient()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func CountPerHospital(c *gin.Context) {
	//call pugin
	data, er := model.CountPerHospital()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}

func ListTopHospital(c *gin.Context) {
	//call pugin
	data, er := model.TopHospital()
	if er != nil {
		return
	}

	c.JSON(http.StatusOK, data)
	return
}
