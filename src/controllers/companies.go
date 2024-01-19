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
