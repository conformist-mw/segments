package admin

import (
	"errors"
	"strconv"

	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type FormError struct {
	Field   string
	Message string
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
	c.HTML(200, "admin/index.html", gin.H{})
}

func Users(c *gin.Context) {
	form := models.CreateUserForm{}
	c.HTML(200, "admin/users.html", gin.H{"Users": models.GetUsers(), "Form": form})
}

func CreateUser(c *gin.Context) {
	var form models.CreateUserForm
	c.Bind(&form)
	err := models.ValidateCreateUserForm(form)
	if err != nil {
		c.HTML(400, "admin/users.html", gin.H{"User": c.Keys["CurrentUser"], "Error": err, "Form": form, "Users": models.GetUsers()})
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
	c.HTML(200, "admin/color_types.html", gin.H{"ColorTypes": models.GetColorTypes(), "Form": models.ColorTypeForm{}})
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
		c.HTML(400, "admin/color_types.html", gin.H{"Errors": errs, "Form": form, "ColorTypes": models.GetColorTypes()})
		return
	}
	_, err := models.CreateColorType(form)
	if err != nil {
		c.HTML(400, "admin/color_types.html", gin.H{"Error": err.Error(), "Form": form, "ColorTypes": models.GetColorTypes()})
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

func GetColors(c *gin.Context) {
	c.HTML(200, "admin/colors.html", gin.H{"Colors": models.GetColors()})
}

func GetSegments(c *gin.Context) {
	c.HTML(200, "admin/segments.html", gin.H{
		"Segments":  models.GetAdminSegments(),
		"Companies": models.GetCompanies(),
		"Sections":  models.GetAdminSections(),
		"Racks":     models.GetAdminRacks(),
	})
}
