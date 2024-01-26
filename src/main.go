package main

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"

	"github.com/conformist-mw/segments/controllers"
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

var location *time.Location

const userkey = "user"

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

func replace(input, from string, to int) string {
	return strings.ReplaceAll(input, from, strconv.Itoa(to))
}

func add(a, b int) int {
	return a + b
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Keys["User"] = user
	c.Next()
}

func main() {
	models.ConnectDb()
	store := gormsessions.NewStore(models.DB, true, []byte("secret"))

	router := gin.Default()
	router.Use(sessions.Sessions("session", store))
	router.SetFuncMap(template.FuncMap{
		"formatInLocation": FormatInLocation,
		"dict":             dict,
		"seq":              seq,
		"replace":          replace,
		"add":              add,
	})
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")

	auth := router.Group("/")
	auth.Use(AuthRequired)
	{
		auth.GET("/", controllers.GetCompanies)
		auth.GET("/:company", controllers.GetSections)
		auth.GET("/:company/:section", controllers.GetSegments)
		auth.POST("/:company/:section/add", controllers.AddSegment)
		auth.POST("/:company/:section/print", controllers.PrintSegments)
		auth.POST("/:company/:section/move/:segment_id", controllers.MoveSegment)
		auth.POST("/:company/:section/activate/:segment_id", controllers.ActivateSegment)
		auth.POST("/:company/:section/remove/:segment_id", controllers.RemoveSegment)

	}

	router.GET("/login", controllers.LoginForm)
	router.POST("/login", controllers.Login)
	router.Run()
}
