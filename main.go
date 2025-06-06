package main

import (
	"errors"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"text/template"
	"time"
	_ "time/tzdata"

	"github.com/gin-contrib/sessions"
	gormsessions "github.com/gin-contrib/sessions/gorm"
	"github.com/go-playground/validator/v10"

	"github.com/conformist-mw/segments/admin"
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

var location *time.Location
var validate *validator.Validate

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

func SuperuserRequired(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get(userkey)
	if username == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}

	user := models.GetUser(username.(string))
	if user.ID == 0 {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	if !user.IsSuperuser {
		c.Redirect(http.StatusFound, "/")
		return
	}
	c.Next()
}

func AuthRequired(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.Redirect(http.StatusFound, "/login")
		return
	}
	c.Keys["User"] = models.GetUser(user.(string))
	c.Next()
}

func UsersAdminRequired(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatus(400)
		return
	}
	user := models.GetUserById(uint(userId))
	if user.ID == 0 {
		c.AbortWithStatus(400)
		return
	}
	c.Keys["CurrentUser"] = user
	c.Next()
}

func validateSlug(fl validator.FieldLevel) bool {
	matched, _ := regexp.MatchString(`^[a-z0-9-_]+$`, fl.Field().String())
	return matched
}

func main() {
	validate = validator.New()
	validate.RegisterValidation("slug", validateSlug)

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
	router.LoadHTMLGlob("templates/**/*")

	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("slug", validateSlug)
	}

	auth := router.Group("/")
	auth.Use(AuthRequired)
	{
		auth.GET("/", GetCompanies)
		auth.POST("/logout", Logout)
		companiesRouter := auth.Group("/companies")
		{
			companiesRouter.GET("/:company", GetSections)
			companiesRouter.GET("/:company/:section", GetSegments)
			companiesRouter.POST("/:company/:section/add", AddSegment)
			companiesRouter.POST("/:company/:section/print", PrintSegments)
			companiesRouter.POST("/:company/:section/move/:segment_id", MoveSegment)
			companiesRouter.POST("/:company/:section/activate/:segment_id", ActivateSegment)
			companiesRouter.POST("/:company/:section/remove/:segment_id", RemoveSegment)
		}

	}

	adminRouter := router.Group("/admin")
	adminRouter.Use(SuperuserRequired)
	{
		adminRouter.GET("", admin.Index)
		adminRouter.GET("/users", admin.Users)
		adminRouter.POST("/users", admin.CreateUser)
		adminUsersRouter := adminRouter.Group("users")
		adminUsersRouter.Use(UsersAdminRequired)
		{
			adminUsersRouter.GET("/:id", admin.GetUserViewRow)
			adminUsersRouter.GET("/:id/edit", admin.GetUserEditRow)
			adminUsersRouter.PATCH("/:id", admin.UpdateUserRow)
			adminUsersRouter.DELETE("/:id", admin.DeleteUser)
			adminUsersRouter.GET("/:id/change-password", admin.GetUserPasswordRow)
			adminUsersRouter.POST("/:id/change-password", admin.ChangeUserPassword)
		}

		adminRouter.GET("/color-types", admin.GetColorTypes)
		colorTypeRouter := adminRouter.Group("/color-types")
		{
			colorTypeRouter.POST("", admin.CreateColorType)
			colorTypeRouter.DELETE("/:id", admin.DeleteColorType)
		}
		adminRouter.GET("/colors", admin.GetColors)
		colorsRouter := adminRouter.Group("/colors")
		{
			colorsRouter.POST("", admin.CreateColor)
			colorsRouter.GET("/:id", admin.GetColorViewRow)
			colorsRouter.GET("/:id/edit", admin.GetColorEditRow)
			colorsRouter.PATCH("/:id", admin.UpdateColorRow)
			colorsRouter.DELETE("/:id", admin.DeleteColor)
		}
		adminRouter.GET("/segments", admin.GetSegments)
	}

	router.GET("/login", LoginForm)
	router.POST("/login", Login)
	router.Run()
}
