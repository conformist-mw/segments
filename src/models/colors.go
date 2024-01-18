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
	ID     uint   `gorm:"primarykey;not null" json:"id"`
	Name   string `gorm:"type:varchar(30);not null;uniqueIndex:idx_color_type_name" json:"name"`
	Slug   string `gorm:"type:varchar(45);not null" json:"slug"`
	TypeID uint   `gorm:"foreignkey:ColorTypeID;not null;uniqueIndex:idx_color_type_name" json:"type"`
}

func (Color) TableName() string {
	return "segments_color"
}
