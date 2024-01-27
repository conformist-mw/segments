package main

import (
	"strconv"

	"github.com/conformist-mw/segments/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func render(c *gin.Context, templateName string, data gin.H) {
	if data == nil {
		data = gin.H{}
	}
	data["Username"] = c.Keys["User"]
	c.HTML(200, templateName, data)
}

func GetCompanies(c *gin.Context) {
	render(c, "companies.html", gin.H{"Companies": models.GetCompanies()})
}

func GetSections(c *gin.Context) {
	render(c, "sections.html", gin.H{
		"Sections": models.GetSections(c.Param("company")),
		"Company":  models.GetCompany(c.Param("company")),
	})
}

func GetSegments(c *gin.Context) {
	var SearchForm models.SearchForm
	c.Bind(&SearchForm)
	section := models.GetSection(c.Param("section"))
	segments, paginator := models.GetSegments(c.Param("section"), c.Param("company"), SearchForm)
	render(c, "segments.html", gin.H{
		"Segments":   segments,
		"Section":    section,
		"Company":    models.GetCompany(c.Param("company")),
		"Racks":      section.Racks,
		"Colors":     models.GetColors(),
		"ColorTypes": models.GetColorTypes(),
		"SearchForm": SearchForm,
		"Paginator":  paginator,
	})
}

func AddSegment(c *gin.Context) {
	var AddForm models.AddForm
	c.ShouldBind(&AddForm)
	colorErr := models.ValidateColor(AddForm.Color, AddForm.ColorType)
	if colorErr != nil {
		c.JSON(400, gin.H{"error": colorErr.Error()})
		return
	}
	rackErr := models.ValidateRack(c.Param("company"), c.Param("section"), AddForm.RackID)
	if rackErr != nil {
		c.JSON(400, gin.H{"error": rackErr.Error()})
		return
	}
	models.AddSegment(c.Param("section"), c.Param("company"), AddForm)
	c.JSON(201, gin.H{})
}

func MoveSegment(c *gin.Context) {
	var MoveForm models.MoveForm
	c.ShouldBind(&MoveForm)
	rackErr := models.ValidateRack(c.Param("company"), c.Param("section"), MoveForm.Rack)
	if rackErr != nil {
		c.JSON(400, gin.H{"error": rackErr.Error()})
		return
	}
	segmentId, err := strconv.Atoi(c.Param("segment_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.MoveSegment(segmentId, MoveForm)
	c.JSON(200, gin.H{})
}

func ActivateSegment(c *gin.Context) {
	segmentId, err := strconv.Atoi(c.Param("segment_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	models.ActivateSegment(segmentId)
	c.JSON(200, gin.H{})
}

func RemoveSegment(c *gin.Context) {
	// TODO: check if segment belongs to the company
	segmentId, err := strconv.Atoi(c.Param("segment_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var removeForm models.RemoveForm
	c.Bind(&removeForm)
	formErr := models.ValidateRemoveForm(removeForm)
	if formErr != nil {
		c.JSON(400, gin.H{"error": formErr.Error()})
		return
	}
	models.RemoveSegment(segmentId, removeForm)
	c.JSON(200, gin.H{})
}

func PrintSegments(c *gin.Context) {
	var PrintForm models.PrintForm
	c.Bind(&PrintForm)
	segments := models.GetPrintSegments(c.Param("section"), c.Param("company"), PrintForm)
	c.HTML(200, "table.html", gin.H{
		"Segments": segments,
	})
}

func LoginForm(c *gin.Context) {
	c.HTML(200, "login.html", gin.H{})
}

func Login(c *gin.Context) {
	var loginForm models.LoginForm
	c.Bind(&loginForm)
	_, err := models.CheckLogin(loginForm)
	if err != nil {
		c.HTML(400, "login.html", gin.H{"error": err.Error()})
		return
	}
	session := sessions.Default(c)
	session.Set(userkey, loginForm.Username)
	session.Save()
	c.Redirect(302, "/")
}

func Logout(c *gin.Context) {
	session := sessions.Default(c)
	user := session.Get(userkey)
	if user == nil {
		c.Redirect(302, "/login")
		return
	}
	session.Delete(userkey)
	session.Save()
	c.Redirect(302, "/login")
}
