package main

import (
	"errors"
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

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}

func seq(start, end int) []int {
	seq := make([]int, end-start+1)
	for i := range seq {
		seq[i] = start + i
	}
	return seq
}

func main() {
	models.ConnectDb()
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatInLocation": FormatInLocation,
		"dict":             dict,
		"seq":              seq,
	})
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/", controllers.GetCompanies)
	router.GET("/:company", controllers.GetSections)
	router.GET("/:company/:section", controllers.GetSegments)
	router.Run()
}
