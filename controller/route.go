package controller

//get url
import (
	_ "MVC63_1/view"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	//_ "github.com/zxc10110/mvc_12/view"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	//r.POST call function list
	r.GET("/listHospital", ListHospital)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r

}
