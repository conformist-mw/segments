package admin

import (
	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.HTML(200, "admin_index.html", gin.H{})
}

func Users(c *gin.Context) {
	form := models.CreateUserForm{}
	c.HTML(200, "admin_users.html", gin.H{"Users": models.GetUsers(), "Form": form})
}

func CreateUser(c *gin.Context) {
	var form models.CreateUserForm
	c.Bind(&form)
	err := models.ValidateCreateUserForm(form)
	if err != nil {
		c.HTML(400, "admin_users.html", gin.H{"User": c.Keys["CurrentUser"], "Error": err, "Form": form, "Users": models.GetUsers()})
		return
	}
	models.CreateUser(form)
	c.Redirect(302, "/admin/users")
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
