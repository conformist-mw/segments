package controllers

import (
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
	section := models.GetSection(c.Param("section"))
	c.HTML(200, "segments.html", gin.H{
		"Segments":   models.GetSegments(c.Param("section"), c.Param("company")),
		"Section":    section,
		"Company":    models.GetCompany(c.Param("company")),
		"Racks":      section.Racks,
		"Colors":     models.GetColors(),
		"ColorTypes": models.GetColorTypes(),
	})
}
