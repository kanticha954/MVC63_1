package main

import (
	"fmt"

	"MVC63_1/configure"
	route "MVC63_1/controller"
	"MVC63_1/model"

	"github.com/jinzhu/gorm"
	//"github.com/zxc10110/mvc_/Models"
	//"github.com/zxc10110/mvc_12/Config"
	//Routes "github.com/zxc10110/mvc_12/Controllers"
)

var err error

func main() {
	//swagger
	fmt.Println("Run Swagger : http://localhost:8080/swagger/index.html")
	fmt.Println()
	//database configure
	configure.DB, err = gorm.Open("mysql", configure.DbURL(configure.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
		//fmt.Println("CONNECTED ERROR")
	}
	defer configure.DB.Close()
	configure.DB.AutoMigrate(&model.Hospital{}) //create to student
	r := route.SetupRouter()
	//running
	r.Run()
}
