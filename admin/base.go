package admin

import "github.com/gin-gonic/gin"

func AdminIndex(c *gin.Context) {
	c.HTML(200, "admin_index.html", gin.H{})
}
