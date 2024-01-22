package models

import "time"

type OrderNumber struct {
	ID   uint   `gorm:"primarykey;not null" json:"id"`
	Name string `gorm:"type:varchar(15);not null;unique" json:"name"`
}

func (OrderNumber) TableName() string {
	return "segments_ordernumber"
}

type Segment struct {
	ID            uint      `gorm:"primarykey;not null" json:"id"`
	Width         int       `gorm:"type:int;not null" json:"width"`
	Height        int       `gorm:"type:int;not null" json:"height"`
	Square        float64   `gorm:"type:float;not null" json:"square"`
	Description   string    `gorm:"type:text;not null" json:"description"`
	Created       time.Time `gorm:"type:datetime;not null" json:"created"`
	Deleted       time.Time `gorm:"type:datetime" json:"deleted"`
	Defect        bool      `gorm:"type:bool;not null;default:false" json:"defect"`
	Active        bool      `gorm:"type:bool;not null;default:true" json:"active"`
	ColorID       uint      `gorm:"foreignkey:ColorID;references:ID;not null" json:"color_id"`
	OrderNumberID *uint     `gorm:"foreignkey:OrderNumberID;references:ID" json:"order_number_id"`
	RackID        uint      `gorm:"foreignkey:RackID;references:ID;not null" json:"rack_id"`
	OrderNumber   OrderNumber
	Color         Color
	Rack          Rack
}

func (Segment) TableName() string {
	return "segments_segment"
}

func GetSegments(sectionSlug, companySlug string) []Segment {
	var segments []Segment

	DB.Preload("Color").
		Preload("Color.Type").
		Preload("OrderNumber").
		Preload("Rack").
		Joins("JOIN segments_rack on segments_rack.id = segments_segment.rack_id").
		Joins("JOIN segments_section on segments_section.id = segments_rack.section_id").
		Joins("JOIN segments_company on segments_company.id = segments_section.company_id").
		Where("segments_section.slug = ? AND segments_company.slug = ?", sectionSlug, companySlug).
		Order("created desc").
		Find(&segments)

	return segments
}