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

func GetUserEditForm(c *gin.Context) {
	c.HTML(200, "admin_user_edit_form.html", gin.H{"User": models.GetUser(c.Param("username"))})
}
