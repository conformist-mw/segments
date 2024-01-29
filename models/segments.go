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

func GetOrderNumber(orderNumber string) OrderNumber {
	var order OrderNumber
	DB.Where(&OrderNumber{Name: orderNumber}).First(&order)
	return order
}

func GetSegment(segmentId int) Segment {
	var segment Segment
	DB.Where(&Segment{ID: uint(segmentId)}).First(&segment)
	return segment
}

func GetSegments(sectionSlug string, companySlug string, SearchForm SearchForm) ([]Segment, Paginator) {
	var segments []Segment
	var rowsCount int64
	var paginator Paginator

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
	countDB := *db
	countDB.Model(&Segment{}).Count(&rowsCount)
	paginator.TotalItems = rowsCount
	paginator.Current = SearchForm.GetPage()
	paginator.SearchForm = SearchForm
	db.Offset(paginator.GetOffset()).Limit(paginator.GetLimit()).Find(&segments)
	return segments, paginator
}

func AddSegment(sectionSlug string, companySlug string, AddForm AddForm) {
	var segment Segment
	color := GetColor(AddForm.Color)

	segment.Width = AddForm.Width
	segment.Height = AddForm.Height
	segment.Square = float64(AddForm.Width*AddForm.Height) / 10000
	segment.ColorID = color.ID
	segment.RackID = uint(AddForm.RackID)

	DB.Create(&segment)
}

func MoveSegment(segmentId int, MoveForm MoveForm) {
	var segment Segment
	segment = GetSegment(segmentId)
	segment.RackID = uint(MoveForm.Rack)
	DB.Save(&segment)
}

func ActivateSegment(segmentId int) {
	var segment Segment
	segment = GetSegment(segmentId)
	segment.Active = true
	segment.OrderNumberID = nil
	segment.Description = ""
	DB.Save(&segment)
}

func RemoveSegment(segmentId int, RemoveForm RemoveForm) {
	var segment Segment
	segment = GetSegment(segmentId)
	segment.Active = false
	if RemoveForm.OrderNumber != "" {
		var orderNumber OrderNumber
		orderNumber.Name = RemoveForm.OrderNumber
		DB.Create(&orderNumber)
		segment.OrderNumberID = &orderNumber.ID
	}
	if RemoveForm.Defect == "on" {
		segment.Defect = true
		segment.Description = RemoveForm.Description
	}
	DB.Save(&segment)
}

func GetPrintSegments(sectionSlug string, companySlug string, PrintForm PrintForm) []Segment {
	var segments []Segment
	db := DB.Preload("Color").
		Preload("Color.Type").
		Preload("Rack").
		Joins("JOIN segments_rack on segments_rack.id = segments_segment.rack_id").
		Joins("JOIN segments_section on segments_section.id = segments_rack.section_id").
		Joins("JOIN segments_company on segments_company.id = segments_section.company_id").
		Where("segments_section.slug = ? AND segments_company.slug = ?", sectionSlug, companySlug).
		Order("created desc")

	if PrintForm.PrintRack != 0 {
		db = db.Where("segments_rack.id = ?", PrintForm.PrintRack)
	}
	db.Find(&segments)
	return segments
}
