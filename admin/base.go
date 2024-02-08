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
	c.HTML(200, "admin_user_row.html", gin.H{"User": c.Keys["CurrentUser"]})
}

func UpdateUserRow(c *gin.Context) {
	var form models.UserUpdateForm
	var user models.User
	c.Bind(&form)
	user = c.Keys["CurrentUser"].(models.User)
	user = models.UpdateUser(user, form)
	c.HTML(200, "admin_user_row.html", gin.H{"User": user})
}

func GetUserPasswordRow(c *gin.Context) {
	c.HTML(200, "admin_user_change_password_row.html", gin.H{"User": c.Keys["CurrentUser"]})
}

func ChangeUserPassword(c *gin.Context) {
	var form models.ChangePasswordForm
	var user models.User
	c.ShouldBind(&form)
	user = c.Keys["CurrentUser"].(models.User)
	user = models.ChangePassword(user, form)
	c.HTML(200, "admin_user_row.html", gin.H{"User": user})
}

func DeleteUser(c *gin.Context) {
	user := c.Keys["CurrentUser"].(models.User)
	models.DeleteUser(user)
	c.Status(200)
}
