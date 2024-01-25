package controllers

import (
	"strconv"

	"github.com/conformist-mw/segments/models"
	"github.com/gin-gonic/gin"
)

func GetCompanies(c *gin.Context) {
	c.HTML(200, "companies.html", gin.H{"Companies": models.GetCompanies()})
}

func GetSections(c *gin.Context) {

	c.HTML(200, "sections.html", gin.H{
		"Sections": models.GetSections(c.Param("company")),
		"Company":  models.GetCompany(c.Param("company")),
	})
}

func GetSegments(c *gin.Context) {
	var SearchForm models.SearchForm
	c.Bind(&SearchForm)
	section := models.GetSection(c.Param("section"))
	segments, paginator := models.GetSegments(c.Param("section"), c.Param("company"), SearchForm)
	c.HTML(200, "segments.html", gin.H{
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
