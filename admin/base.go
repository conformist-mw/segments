package admin

import (
	"strconv"

	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

func getUser(c *gin.Context) models.User {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return models.User{}
	}
	user := models.GetUserById(uint(userId))
	if user.ID == 0 {
		return models.User{}
	}
	return user
}

func Index(c *gin.Context) {
	c.HTML(200, "admin_index.html", gin.H{})
}

func Users(c *gin.Context) {
	c.HTML(200, "admin_users.html", gin.H{"Users": models.GetUsers()})
}

func GetUserEditRow(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "Bad request")
		return
	}
	c.HTML(200, "admin_user_edit_row.html", gin.H{"User": models.GetUserById(uint(userId))})
}

func GetUserViewRow(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.String(400, "Bad request")
		return
	}
	c.HTML(200, "admin_user_view_row.html", gin.H{"User": models.GetUserById(uint(userId))})
}
