package models

import (
	"math"
	"net/url"
	"strconv"
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

func (s SearchForm) GetPage() int {
	if s.Page <= 0 {
		return 1
	}
	return s.Page
}

type Paginator struct {
	TotalItems int64
	Current    int
	Next       int
	Previous   int
	SearchForm SearchForm
}

func (p Paginator) GetOffset() int {
	return 10 * (p.Current - 1)
}

func (p Paginator) GetLimit() int {
	return 10
}

func (p Paginator) GetTotalPages() int {
	return int(math.Ceil(float64(p.TotalItems) / float64(p.GetLimit())))
}

func (p Paginator) HasPrevious() bool {
	return p.Current > 1
}

func (p Paginator) GetPrevious() int {
	if p.HasPrevious() {
		return p.Current - 1
	}
	return p.Current
}

func (p Paginator) HasNext() bool {
	return p.Current < p.GetTotalPages()
}

func (p Paginator) GetNext() int {
	if p.HasNext() {
		return p.Current + 1
	}
	return p.Current
}

func (p Paginator) Start() int {
	start := p.Current - 6
	if start < 1 {
		start = 1
	}
	return start
}

func (p Paginator) End() int {
	end := p.Current + 6
	if end > p.GetTotalPages() {
		end = p.GetTotalPages()
	}
	return end
}

func (p Paginator) PageUrl() string {
	params := make(url.Values)
	params.Set("color", p.SearchForm.Color)
	params.Set("color_type", p.SearchForm.ColorType)
	params.Set("deleted", p.SearchForm.Deleted)
	params.Set("height", strconv.Itoa(p.SearchForm.Height))
	params.Set("width", strconv.Itoa(p.SearchForm.Width))
	params.Set("page", "__page_number__")
	return "?" + params.Encode()
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
