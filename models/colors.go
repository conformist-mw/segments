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
