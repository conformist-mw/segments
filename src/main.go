package main

import (
	"os"
	"text/template"
	"time"

	"github.com/conformist-mw/segments/controllers"
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

var location *time.Location

func init() {
	os.Setenv("TZ", "Europe/Kyiv")
	var err error
	location, err = time.LoadLocation("Europe/Kyiv")
	if err != nil {
		panic(err)
	}
}

func FormatInLocation(t time.Time) string {
	return t.In(location).Format("02.01.06 | 15:04")
}

func main() {
	models.ConnectDb()
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatInLocation": FormatInLocation,
	})
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.GetCompanies)
	router.GET("/:company", controllers.GetSections)
	router.GET("/:company/:section", controllers.GetSegments)
	router.Run()
}
