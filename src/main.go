package main

import (
	"github.com/conformist-mw/segments/controllers"
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

func main() {
	models.ConnectDb()
	router := gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.GetCompanies)
	router.GET("/:company", controllers.GetSections)
	router.Run()
}
