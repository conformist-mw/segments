package models

import (
	"errors"
	"fmt"
)

type ColorType struct {
	ID     uint    `gorm:"primarykey;not null" json:"id"`
	Name   string  `gorm:"type:varchar(15);unique;not null" json:"name"`
	Slug   string  `gorm:"type:varchar(25);unique;not null" json:"slug"`
	Colors []Color `gorm:"foreignKey:TypeID" json:"colors"`
}

func (ColorType) TableName() string {
	return "segments_colortype"
}

type Color struct {
	ID     uint      `gorm:"primarykey;not null" json:"id"`
	Name   string    `gorm:"type:varchar(30);not null;uniqueIndex:idx_color_type_name_uniq" json:"name"`
	Slug   string    `gorm:"type:varchar(45);not null" json:"slug"`
	TypeID uint      `gorm:"foreignkey:TypeID;references:ID;not null;uniqueIndex:idx_color_type_name_uniq" json:"type_id"`
	Type   ColorType `gorm:"foreignkey:TypeID;references:ID;not null" json:"type"`
}

func (Color) TableName() string {
	return "segments_color"
}

func GetColors() []Color {
	var colors []Color
	DB.Preload("Type").Find(&colors)
	return colors
}

func GetColor(slug string) Color {
	var color Color
	DB.Preload("Type").Where(&Color{Slug: slug}).First(&color)
	return color
}

func GetColorTypes() []ColorType {
	var colorTypes []ColorType
	DB.Find(&colorTypes)
	return colorTypes
}

func GetColorType(slug string) ColorType {
	var colorType ColorType
	DB.Where(&ColorType{Slug: slug}).First(&colorType)
	return colorType
}

type ColorTypeForm struct {
	Name string `form:"name" binding:"required"`
	Slug string `form:"slug" binding:"required,slug"`
}

func CreateColorType(form ColorTypeForm) (ColorType, error) {
	colorType := ColorType{
		Name: form.Name,
		Slug: form.Slug,
	}
	result := DB.Create(&colorType)
	if result.Error != nil {
		return ColorType{}, result.Error
	}
	return colorType, nil
}

func UpdateColorType(slug string, form ColorTypeForm) (ColorType, error) {
	colorType := ColorType{
		Name: form.Name,
		Slug: form.Slug,
	}
	result := DB.Where(&ColorType{Slug: slug}).Updates(&colorType)
	if result.Error != nil {
		return ColorType{}, result.Error
	}
	return colorType, nil
}

func DeleteColorType(id uint) error {
	colorType := ColorType{}
	DB.Preload("Colors").Where(&ColorType{ID: id}).First(&colorType)
	if colorType.ID == 0 {
		return fmt.Errorf("Color type with id %d not found", id)
	}
	if len(colorType.Colors) > 0 {
		return errors.New("Color type has colors")
	}
	result := DB.Delete(&colorType)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetColorById(id uint) Color {
	var color Color
	DB.Preload("Type").Where(&Color{ID: id}).First(&color)
	return color
}

func CreateColor(form ColorForm) (Color, error) {
	colorType := GetColorType(form.ColorType)
	if colorType.ID == 0 {
		return Color{}, fmt.Errorf("color type not found")
	}
	color := Color{
		Name:   form.Name,
		Slug:   form.Slug,
		TypeID: colorType.ID,
	}
	result := DB.Create(&color)
	if result.Error != nil {
		return Color{}, result.Error
	}
	DB.Preload("Type").First(&color, color.ID)
	return color, nil
}

func UpdateColor(id uint, form ColorForm) (Color, error) {
	color := GetColorById(id)
	if color.ID == 0 {
		return Color{}, fmt.Errorf("color not found")
	}
	if form.Name != "" {
		color.Name = form.Name
	}
	if form.Slug != "" {
		color.Slug = form.Slug
	}
	if form.ColorType != "" {
		colorType := GetColorType(form.ColorType)
		if colorType.ID == 0 {
			return Color{}, fmt.Errorf("color type not found")
		}
		color.TypeID = colorType.ID
	}
	result := DB.Save(&color)
	if result.Error != nil {
		return Color{}, result.Error
	}
	DB.Preload("Type").First(&color, color.ID)
	return color, nil
}

func DeleteColor(id uint) error {
	var count int64
	DB.Model(&Segment{}).Where(&Segment{ColorID: id}).Count(&count)
	if count > 0 {
		return errors.New("color has segments")
	}
	result := DB.Delete(&Color{}, id)
	return result.Error
}
