package models

import (
	"time"
)

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

type SearchForm struct {
	ColorType   string `form:"color_type"`
	Color       string `form:"color"`
	Width       int    `form:"width"`
	Height      int    `form:"height"`
	OrderNumber string `form:"order_number"`
	Deleted     string `form:"deleted"`
	Page        int    `form:"page"`
}

// if width := fields['width']:
// 	qs = qs.filter(Q(width__gte=width) | Q(height__gte=width))
// if height := fields['height']:
// 	qs = qs.filter(Q(width__gte=height) | Q(height__gte=height))
// active = not fields.get('deleted', False)
// qs = qs.filter(active=active)
// if not active and (order_number := fields.get('order_number')):
// 	qs = qs.filter(order_number__name=order_number)

func GetSegments(sectionSlug string, companySlug string, SearchForm SearchForm) []Segment {
	var segments []Segment

	db := DB.Preload("Color").
		Preload("Color.Type").
		Preload("OrderNumber").
		Preload("Rack").
		Joins("JOIN segments_rack on segments_rack.id = segments_segment.rack_id").
		Joins("JOIN segments_section on segments_section.id = segments_rack.section_id").
		Joins("JOIN segments_company on segments_company.id = segments_section.company_id").
		Joins("JOIN segments_color on segments_color.id = segments_segment.color_id").
		Joins("JOIN segments_colortype on segments_colortype.id = segments_color.type_id").
		Joins("LEFT JOIN segments_ordernumber on segments_ordernumber.id = segments_segment.order_number_id").
		Where("segments_section.slug = ? AND segments_company.slug = ?", sectionSlug, companySlug).
		Order("created desc")

	if SearchForm.ColorType != "" {
		db = db.Where("segments_colortype.slug = ?", SearchForm.ColorType)
	}
	if SearchForm.Color != "" {
		db = db.Where("segments_color.slug = ?", SearchForm.Color)
	}
	if SearchForm.Width != 0 {
		db = db.Where("segments_segment.width >= ? OR segments_segment.height >= ?", SearchForm.Width, SearchForm.Width)
	}
	if SearchForm.Height != 0 {
		db = db.Where("segments_segment.width >= ? OR segments_segment.height >= ?", SearchForm.Height, SearchForm.Height)
	}
	if SearchForm.Deleted == "on" {
		db = db.Where("segments_segment.active IS FALSE")
	} else {
		db = db.Where("segments_segment.active IS TRUE")
	}
	if SearchForm.OrderNumber != "" {
		db = db.Where("segments_ordernumber.name = ?", SearchForm.OrderNumber)
	}

	db.Find(&segments)
	return segments
}
