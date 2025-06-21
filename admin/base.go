package admin

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FormError struct {
	Field   string
	Message string
}

func renderAdmin(c *gin.Context, templateName string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	data["User"] = c.Keys["User"]
	if contextPage, exists := c.Get("CurrentPage"); exists {
		data["CurrentPage"] = contextPage
	}
	fmt.Printf("Rendering admin page: %s, CurrentPage: %s\n", templateName, data["CurrentPage"])
	c.HTML(200, templateName, data)
}

func renderAdminError(c *gin.Context, templateName string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	data["User"] = c.Keys["User"]
	if contextPage, exists := c.Get("CurrentPage"); exists {
		data["CurrentPage"] = contextPage
	}
	c.HTML(400, templateName, data)
}

func GetFormErrors(form interface{}, c *gin.Context) (errs []FormError) {
	if formErr := c.ShouldBind(form); formErr != nil {
		var ve validator.ValidationErrors
		if errors.As(formErr, &ve) {
			for _, e := range ve {
				errs = append(errs, FormError{Field: e.Field(), Message: MessageForTag(e.Tag())})
			}
		}
		return errs
	}
	return nil
}

func GetUintId(c *gin.Context) (uint, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return 0, err
	}
	return uint(id), nil
}

func Index(c *gin.Context) {
	renderAdmin(c, "admin/index.html", gin.H{})
}

func Users(c *gin.Context) {
	form := models.CreateUserForm{}
	renderAdmin(c, "admin/users.html", gin.H{"Users": models.GetUsers(), "Form": form})
}

func CreateUser(c *gin.Context) {
	var form models.CreateUserForm
	c.Bind(&form)
	err := models.ValidateCreateUserForm(form)
	if err != nil {
		renderAdminError(c, "admin/users.html", gin.H{"User": c.Keys["CurrentUser"], "Error": err, "Form": form, "Users": models.GetUsers()})
		return
	}
	models.CreateUser(form)
	c.Redirect(302, "/admin/users")
}

func GetUserEditRow(c *gin.Context) {
	c.HTML(200, "admin_user_edit_row", c.Keys["CurrentUser"])
}

func GetUserViewRow(c *gin.Context) {
	c.HTML(200, "admin_user_row", c.Keys["CurrentUser"])
}

func UpdateUserRow(c *gin.Context) {
	var form models.UserUpdateForm
	var user models.User
	c.Bind(&form)
	user = c.Keys["CurrentUser"].(models.User)
	user = models.UpdateUser(user, form)
	c.HTML(200, "admin_user_row", user)
}

func GetUserPasswordRow(c *gin.Context) {
	c.HTML(200, "admin_user_change_password_row", c.Keys["CurrentUser"])
}

func ChangeUserPassword(c *gin.Context) {
	var form models.ChangePasswordForm
	var user models.User
	c.ShouldBind(&form)
	user = c.Keys["CurrentUser"].(models.User)
	user = models.ChangePassword(user, form)
	c.HTML(200, "admin_user_row", user)
}

func DeleteUser(c *gin.Context) {
	user := c.Keys["CurrentUser"].(models.User)
	models.DeleteUser(user)
	c.Status(200)
}

func GetColorTypes(c *gin.Context) {
	renderAdmin(c, "admin/color_types.html", gin.H{"ColorTypes": models.GetColorTypes(), "Form": models.ColorTypeForm{}})
}

func MessageForTag(tag string) string {
	switch tag {
	case "required":
		return "This field is required"
	case "slug":
		return "Is invalid (only letters, numbers, and dashes)"
	default:
		return "This field is invalid"
	}
}

func CreateColorType(c *gin.Context) {
	var form models.ColorTypeForm

	if errs := GetFormErrors(&form, c); errs != nil {
		renderAdminError(c, "admin/color_types.html", gin.H{"Errors": errs, "Form": form, "ColorTypes": models.GetColorTypes()})
		return
	}
	_, err := models.CreateColorType(form)
	if err != nil {
		renderAdminError(c, "admin/color_types.html", gin.H{"Error": err.Error(), "Form": form, "ColorTypes": models.GetColorTypes()})
		return
	}
	c.Redirect(302, "/admin/color-types")
}

func DeleteColorType(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	errs := models.DeleteColorType(id)
	if errs != nil {
		c.Status(400)
		return
	}
	c.Status(200)
}

func GetColorTypeEditRow(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	colorType := models.GetColorTypeById(id)
	if colorType.ID == 0 {
		c.Status(400)
		return
	}
	c.HTML(200, "admin_color_type_edit_row", colorType)
}

func UpdateColorTypeRow(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	var form models.ColorTypeForm
	c.Bind(&form)
	colorType, err := models.UpdateColorType(id, form)
	if err != nil {
		c.Status(400)
		return
	}
	c.HTML(200, "admin_color_type_row", colorType)
}

func GetColors(c *gin.Context) {
	renderAdmin(c, "admin/colors.html", gin.H{
		"Colors":     models.GetColors(),
		"ColorTypes": models.GetColorTypes(),
		"Form":       models.ColorForm{},
	})
}

func CreateColor(c *gin.Context) {
	var form models.ColorForm
	if errs := GetFormErrors(&form, c); errs != nil {
		renderAdminError(c, "admin/colors.html", gin.H{
			"Errors":     errs,
			"Form":       form,
			"Colors":     models.GetColors(),
			"ColorTypes": models.GetColorTypes(),
		})
		return
	}
	_, err := models.CreateColor(form)
	if err != nil {
		renderAdminError(c, "admin/colors.html", gin.H{
			"Error":      err.Error(),
			"Form":       form,
			"Colors":     models.GetColors(),
			"ColorTypes": models.GetColorTypes(),
		})
		return
	}
	c.Redirect(302, "/admin/colors")
}

func GetColorEditRow(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	color := models.GetColorById(id)
	if color.ID == 0 {
		c.Status(400)
		return
	}
	c.HTML(200, "admin_color_edit_row", gin.H{
		"Color":      color,
		"ColorTypes": models.GetColorTypes(),
	})
}

func GetColorViewRow(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	color := models.GetColorById(id)
	if color.ID == 0 {
		c.Status(400)
		return
	}
	c.HTML(200, "admin_color_row", color)
}

func UpdateColorRow(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	var form models.ColorForm
	c.Bind(&form)
	color, err := models.UpdateColor(id, form)
	if err != nil {
		c.Status(400)
		return
	}
	c.HTML(200, "admin_color_row", color)
}

func DeleteColor(c *gin.Context) {
	id, err := GetUintId(c)
	if err != nil {
		c.Status(400)
		return
	}
	err = models.DeleteColor(id)
	if err != nil {
		c.Status(400)
		return
	}
	c.Status(200)
}

func GetSegments(c *gin.Context) {
	renderAdmin(c, "admin/segments.html", gin.H{
		"Segments":  models.GetAdminSegments(),
		"Companies": models.GetCompanies(),
		"Sections":  models.GetAdminSections(),
		"Racks":     models.GetAdminRacks(),
	})
}
