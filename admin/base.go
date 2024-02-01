package admin

import (
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "admin_index.html", gin.H{})
}

func Users(c *gin.Context) {
	c.HTML(200, "admin_users.html", gin.H{"Users": models.GetUsers()})
}

func GetUserEditRow(c *gin.Context) {
	c.HTML(200, "admin_user_edit_row.html", gin.H{"User": c.Keys["CurrentUser"]})
}

func GetUserViewRow(c *gin.Context) {
	c.HTML(200, "admin_user_view_row.html", gin.H{"User": c.Keys["CurrentUser"]})
}

func UpdateUserRow(c *gin.Context) {
	var form models.UserUpdateForm
	var user models.User
	c.Bind(&form)
	user = c.Keys["CurrentUser"].(models.User)
	user = models.UpdateUser(user, form)
	c.HTML(200, "admin_user_view_row.html", gin.H{"User": user})
}
