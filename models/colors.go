package models

type ColorType struct {
	ID   uint   `gorm:"primarykey;not null" json:"id"`
	Name string `gorm:"type:varchar(15);unique;not null" json:"name"`
	Slug string `gorm:"type:varchar(25);unique;not null" json:"slug"`
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