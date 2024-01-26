package models

import (
	"errors"
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

func GetOrderNumber(orderNumber string) OrderNumber {
	var order OrderNumber
	DB.Where(&OrderNumber{Name: orderNumber}).First(&order)
	return order
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

type AddForm struct {
	Width     int    `form:"width" binding:"required"`
	Height    int    `form:"height" binding:"required"`
	Color     string `form:"color" binding:"required"`
	ColorType string `form:"color_type" binding:"required"`
	RackID    int    `form:"rack_id" binding:"required"`
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

func ValidateColor(colorSlug, colorTypeSlug string) error {
	var color Color
	var colorType ColorType
	DB.Where("slug = ?", colorSlug).First(&color)
	DB.Where("slug = ?", colorTypeSlug).First(&colorType)
	if color.ID == 0 {
		return errors.New("color not found")
	}
	if colorType.ID == 0 {
		return errors.New("color type not found")
	}
	if color.TypeID != colorType.ID {
		return errors.New("color type mismatch")
	}
	return nil
}

func ValidateRack(companySlug, sectionSlug string, rackId int) error {
	rack := GetRack(companySlug, sectionSlug, rackId)
	if rack.ID == 0 {
		return errors.New("rack not found")
	}
	return nil
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

type MoveForm struct {
	Rack int `form:"rack" binding:"required"`
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

type RemoveForm struct {
	OrderNumber string `form:"order_number"`
	Defect      string `form:"defect"`
	Description string `form:"description"`
}

func ValidateRemoveForm(RemoveForm RemoveForm) error {
	isDefect := RemoveForm.Defect == "on"
	if RemoveForm.OrderNumber == "" && !isDefect {
		return errors.New("Для удаления отрезка нужно указать или номер заказа или указать дефект")
	}
	if RemoveForm.OrderNumber == "" {
		if isDefect && RemoveForm.Description == "" {
			return errors.New("При указании дефекта нужно указать описание")
		}
	}
	if RemoveForm.OrderNumber != "" && GetOrderNumber(RemoveForm.OrderNumber).ID != 0 {
		// TODO: order number should be unique within company
		return errors.New("Такой номер заказа уже есть в базе")
	}
	if RemoveForm.OrderNumber == "" && !isDefect {
		return errors.New("no order number or defect")
	}
	return nil
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
